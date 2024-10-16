package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"golang.org/x/tools/go/packages"
	"gopkg.in/yaml.v2"
)

// NewDependencyAnalyzer crea una nueva instancia de DependencyAnalyzer
func NewDependencyAnalyzer() *DependencyAnalyzer {
	return &DependencyAnalyzer{}
}

// AddFileImport agrega un archivo con sus entidades analizadas al DependencyAnalyzer
func (da *DependencyAnalyzer) AddFileImport(name, path, layer string, entities []EntityInfo) {
	da.PackagesInfo = append(da.PackagesInfo, FileImport{
		Name:     name,
		Path:     path,
		Layer:    layer,
		Entities: entities,
	})
}

// ---------------------------
// Carga de configuración de capas desde YAML
// ---------------------------

// loadLayerConfig carga la configuración de capas desde un archivo YAML
func loadLayerConfig(path string) (LayerConfig, error) {
	var config LayerConfig
	content, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(content, &config)
	return config, err
}

// ---------------------------
// Manejo de archivos en el repositorio con go-git
// ---------------------------

// getFilesToAnalyze devuelve una lista de archivos Go para analizar en un repositorio
func getFilesToAnalyze(repo *git.Repository) ([]string, error) {
	var filesToAnalyze []string

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(worktree.Filesystem.Root(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			relPath, err := filepath.Rel(worktree.Filesystem.Root(), path)
			if err != nil {
				return err
			}
			filesToAnalyze = append(filesToAnalyze, relPath)
		}
		return nil
	})

	return filesToAnalyze, err
}

// getFileAuthor obtiene el autor del archivo en el repositorio git
func getFileAuthor(repo *git.Repository, file string) (string, error) {
	commits, err := repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	commit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return commit.Author.Email, nil
}

// getCommitID obtiene el ID del commit para un archivo en el repositorio git
func getCommitID(repo *git.Repository, file string) (string, error) {
	commits, err := repo.Log(&git.LogOptions{FileName: &file})
	if err != nil {
		return "", err
	}

	objectCommit, err := commits.Next()
	if err != nil {
		return "", err
	}

	return objectCommit.Hash.String(), nil
}

// ---------------------------
// Clasificación y análisis de archivos
// ---------------------------

// classifyFile clasifica y analiza un archivo Go
func classifyFile(filePath string, layer string, config LayerConfig, analyzer *DependencyAnalyzer) map[string]SkillData {
	result := make(map[string]SkillData)
	for _, skill := range skills {
		result[skill.ID] = SkillData{Score: 0, Evidence: []Evidence{}}
	}

	entities, _ := listVariablesStructsParamsAndInterfaces(filePath, config)
	pkgName, _ := getPackageName(filePath)
	analyzer.AddFileImport(pkgName, filePath, layer, entities)

	getResults(analyzer, result)

	return result
}

// listVariablesStructsParamsAndInterfaces extrae entidades (structs, variables, etc.) de un archivo Go
func listVariablesStructsParamsAndInterfaces(filePath string, config LayerConfig) ([]EntityInfo, error) {
	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo, Dir: filepath.Dir(filePath)}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil || len(pkgs) == 0 {
		return nil, err
	}

	var results []EntityInfo
	pkg := pkgs[0]
	file, err := getFileFromPackage(pkg, filePath)
	if err != nil {
		return nil, err
	}

	imports := extractImports(file)

	ast.Inspect(file, func(n ast.Node) bool {
		switch decl := n.(type) {
		case *ast.GenDecl:
			processVariables(decl, pkg, imports, config, &results)
		case *ast.FuncDecl:
			processFunctionParamsAndResults(decl, pkg, imports, config, &results)
		case *ast.TypeSpec:
			processStructFields(decl, pkg, imports, config, &results)
		}
		return true
	})

	return results, nil
}

// ---------------------------
// Procesamiento de variables, funciones y structs
// ---------------------------

// processStructFields procesa las entidades de tipo struct en un archivo Go
func processStructFields(typeDecl *ast.TypeSpec, pkg *packages.Package, imports map[string]string, config LayerConfig, results *[]EntityInfo) {
	if structType, ok := typeDecl.Type.(*ast.StructType); ok {
		for _, field := range structType.Fields.List {
			fieldType := getTypeFromAST(field.Type, imports)
			kind := getKindFromType(field.Type, pkg)
			layer := getLayerForType(fieldType, config)

			for _, fieldName := range field.Names {
				line := pkg.Fset.Position(field.Pos()).Line
				*results = append(*results, EntityInfo{
					Name:        fieldName.Name,
					Type:        fieldType,
					Position:    line,
					Category:    "Struct Field",
					Kind:        kind,
					IsInterface: kind == "interface",
					Layer:       layer,
				})
			}
		}
	}
}

// processVariables procesa las variables globales o locales en un archivo Go
func processVariables(decl *ast.GenDecl, pkg *packages.Package, imports map[string]string, config LayerConfig, results *[]EntityInfo) {
	for _, spec := range decl.Specs {
		if vspec, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range vspec.Names {
				varType := getVariableType(vspec.Type, name, pkg, imports)
				category := "Global Variable"
				if !isGlobalVariable(pkg.Fset.Position(name.Pos()).Line, pkg) {
					category = "Local Variable"
				}
				layer := getLayerForType(varType, config)
				kind := getKindFromObj(pkg.TypesInfo.ObjectOf(name))

				line := pkg.Fset.Position(name.Pos()).Line
				*results = append(*results, EntityInfo{
					Name:        name.Name,
					Type:        varType,
					Position:    line,
					Category:    category,
					Kind:        kind,
					IsInterface: kind == "interface",
					Layer:       layer,
				})
			}
		}
	}
}

// processFunctionParamsAndResults procesa los parámetros y resultados de una función Go
func processFunctionParamsAndResults(funcDecl *ast.FuncDecl, pkg *packages.Package, imports map[string]string, config LayerConfig, results *[]EntityInfo) {
	processParamsOrResults(funcDecl.Type.Params, "Function Parameter", pkg, imports, config, results)
	processParamsOrResults(funcDecl.Type.Results, "Function Return Type", pkg, imports, config, results)

	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if assignStmt, ok := n.(*ast.AssignStmt); ok {
			for _, lhs := range assignStmt.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					obj := pkg.TypesInfo.ObjectOf(ident)
					if obj != nil {
						kind := getKindFromObj(obj)
						varType := obj.Type().String()
						layer := getLayerForType(varType, config)
						line := pkg.Fset.Position(ident.Pos()).Line
						*results = append(*results, EntityInfo{
							Name:        ident.Name,
							Type:        varType,
							Position:    line,
							Category:    "Local Variable",
							Kind:        kind,
							IsInterface: kind == "interface",
							Layer:       layer,
						})
					}
				}
			}
		}
		return true
	})
}

// processParamsOrResults procesa los parámetros o resultados de una función Go
func processParamsOrResults(fields *ast.FieldList, category string, pkg *packages.Package, imports map[string]string, config LayerConfig, results *[]EntityInfo) {
	if fields == nil {
		return
	}
	for _, param := range fields.List {
		paramType := getTypeFromAST(param.Type, imports)
		kind := getKindFromType(param.Type, pkg)
		layer := getLayerForType(paramType, config)

		for _, paramName := range param.Names {
			line := pkg.Fset.Position(paramName.Pos()).Line
			*results = append(*results, EntityInfo{
				Name:        paramName.Name,
				Type:        paramType,
				Position:    line,
				Category:    category,
				Kind:        kind,
				IsInterface: kind == "interface",
				Layer:       layer,
			})
		}
	}
}

// ---------------------------
// Utilidades para análisis de tipos y archivos
// ---------------------------

// getVariableType devuelve el tipo de una variable en Go
func getVariableType(expr ast.Expr, name *ast.Ident, pkg *packages.Package, imports map[string]string) string {
	if expr != nil {
		return getTypeFromAST(expr, imports)
	}
	return pkg.TypesInfo.ObjectOf(name).Type().String()
}

// isGlobalVariable determina si una variable es global o local
func isGlobalVariable(line int, pkg *packages.Package) bool {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				start, end := pkg.Fset.Position(funcDecl.Pos()).Line, pkg.Fset.Position(funcDecl.End()).Line
				if line >= start && line <= end {
					return false
				}
			}
		}
	}
	return true
}

// getKindFromObj devuelve el tipo (struct o interface) de una entidad en Go
func getKindFromObj(obj types.Object) string {
	if obj == nil {
		return "unknown"
	}
	switch obj.Type().Underlying().(type) {
	case *types.Interface:
		return "interface"
	case *types.Struct:
		return "struct"
	}
	return "other"
}

// getTypeFromAST devuelve el tipo de una entidad desde su representación AST
func getTypeFromAST(expr ast.Expr, imports map[string]string) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if pkgIdent, ok := t.X.(*ast.Ident); ok {
			return fmt.Sprintf("%s.%s", imports[pkgIdent.Name], t.Sel.Name)
		}
	case *ast.StarExpr:
		return "*" + getTypeFromAST(t.X, imports)
	case *ast.ArrayType:
		return "[]" + getTypeFromAST(t.Elt, imports)
	}
	return "unknown"
}

// ---------------------------
// Incremento del score y manejo de violaciones
// ---------------------------

// incrementScore incrementa el puntaje si se encuentra una violación DIP
func incrementScore(results map[string]SkillData, skillID string, evidence Evidence) {
	skillData := results[skillID]
	skillData.Score += 1
	skillData.Evidence = append(skillData.Evidence, evidence)
	results[skillID] = skillData
}

// getResults procesa los resultados de un análisis de dependencias
func getResults(analyzer *DependencyAnalyzer, results map[string]SkillData) {
	for _, file := range analyzer.PackagesInfo {
		for _, entity := range file.Entities {
			if shouldDisplayEntity(file.Layer, entity) {
				evidence := Evidence{
					File:       file.Path,
					Line:       entity.Position,
					EntityName: entity.Name,
				}
				incrementScore(results, "dip_violation", evidence)
			}
		}
	}
}

// ---------------------------
// Filtrado de entidades según capas
// ---------------------------

// shouldDisplayEntity determina si una entidad debe mostrarse basado en su capa y tipo
func shouldDisplayEntity(fileLayer string, entity EntityInfo) bool {
	if fileLayer == "domain" && entity.Layer == "domain" {
		return false
	}
	if fileLayer == "application" && (entity.Layer == "application" || entity.Layer == "domain") {
		return false
	}
	if fileLayer != "domain" && fileLayer != "application" {
		return false
	}
	if entity.Kind == "interface" || isPrimitiveType(entity.Type) || isStandardLibraryType(entity.Type) {
		return false
	}
	return true
}

// isPrimitiveType determina si un tipo es primitivo
func isPrimitiveType(varType string) bool {
	primitives := []string{"int", "string", "bool", "float32", "float64", "byte", "rune", "complex64", "complex128"}
	for _, primitive := range primitives {
		if varType == primitive {
			return true
		}
	}
	return false
}

// isStandardLibraryType determina si el tipo pertenece a la librería estándar
func isStandardLibraryType(varType string) bool {
	return !strings.Contains(varType, "/")
}

// ---------------------------
// Análisis completo del repositorio
// ---------------------------

// analyzeRepo realiza el análisis del repositorio y devuelve las métricas encontradas
func analyzeRepo(repoPath string) ([]Metric, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return nil, err
	}

	layerConfig, err := loadLayerConfig(filepath.Join(repoPath, "monitor.yml"))
	if err != nil {
		return nil, fmt.Errorf("error loading layer configuration: %v", err)
	}

	filesToAnalyze, err := getFilesToAnalyze(repo)
	if err != nil {
		return nil, err
	}

	results := make(map[string]map[string]SkillData)
	analyzer := NewDependencyAnalyzer()

	for _, filePath := range filesToAnalyze {
		fullPath := filepath.Join(repoPath, filePath)
		author, err := getFileAuthor(repo, filePath)
		if err != nil {
			return nil, err
		}

		commitID, err := getCommitID(repo, filePath)
		if err != nil {
			return nil, err
		}

		layer := determineLayer(filePath, layerConfig)

		if layer == "domain" || layer == "application" {
			fileResults := classifyFile(fullPath, layer, layerConfig, analyzer)
			if _, exists := results[author]; !exists {
				results[author] = make(map[string]SkillData)
				for _, skill := range skills {
					results[author][skill.ID] = SkillData{Score: 0, Evidence: []Evidence{}}
				}
			}

			for skillID, data := range fileResults {
				skillData := results[author][skillID]
				if data.Score > skillData.Score {
					skillData.Score = data.Score
					skillData.Evidence = []Evidence{}
				}
				if data.Score == skillData.Score {
					for _, line := range data.Evidence {
						skillData.Evidence = append(skillData.Evidence, Evidence{
							CommitID:   commitID,
							File:       filePath,
							Line:       line.Line,
							EntityName: line.EntityName,
						})
					}
				}
				results[author][skillID] = skillData
			}
		}
	}

	var output []Metric
	for author, scores := range results {
		for _, skill := range skills {
			skillID := skill.ID
			output = append(output, Metric{
				MetricID:  skillID,
				GitAuthor: author,
				Score:     scores[skillID].Score,
				Evidence:  scores[skillID].Evidence,
			})
		}
	}

	return output, nil
}

// ---------------------------
// Programa principal
// ---------------------------

var skills = []Skill{
	{ID: "dip_violation", Name: "Proper usage of DIP"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <repo_path>")
		return
	}

	repoPath := os.Args[1]
	metrics, err := analyzeRepo(repoPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonOutput, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonOutput))
}

// getKindFromType devuelve el tipo (struct o interface) de una entidad en Go
func getKindFromType(expr ast.Expr, pkg *packages.Package) string {
	if typ, ok := pkg.TypesInfo.Types[expr]; ok {
		switch typ.Type.Underlying().(type) {
		case *types.Interface:
			return "interface"
		case *types.Struct:
			return "struct"
		}
	}
	return "other"
}

// determineLayer determina la capa a la que pertenece un archivo basado en su ruta
func determineLayer(filePath string, config LayerConfig) string {
	for layer, paths := range config.Layers {
		for _, path := range paths {
			if strings.Contains(filePath, path) {
				return layer
			}
		}
	}
	return "other"
}

// getPackageName obtiene el nombre del paquete de un archivo Go
func getPackageName(filePath string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	return node.Name.Name, nil
}

// getFileFromPackage obtiene el archivo Go a partir del paquete
func getFileFromPackage(pkg *packages.Package, filePath string) (*ast.File, error) {
	for _, file := range pkg.Syntax {
		if pkg.Fset.Position(file.Pos()).Filename == filePath {
			return file, nil
		}
	}
	return nil, fmt.Errorf("archivo no encontrado en el paquete")
}

// extractImports extrae las importaciones del archivo Go
func extractImports(file *ast.File) map[string]string {
	imports := make(map[string]string)
	for _, i := range file.Imports {
		alias := ""
		if i.Name != nil {
			alias = i.Name.Name
		} else {
			parts := strings.Split(strings.Trim(i.Path.Value, "\""), "/")
			alias = parts[len(parts)-1]
		}
		imports[alias] = strings.Trim(i.Path.Value, "\"")
	}
	return imports
}

// getLayerForType determina la capa a la que pertenece un tipo específico
func getLayerForType(typeName string, config LayerConfig) string {
	for layer, paths := range config.Layers {
		for _, path := range paths {
			if strings.Contains(typeName, path) {
				return layer
			}
		}
	}
	return "other"
}
