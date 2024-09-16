package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"gopkg.in/yaml.v2"
)

// Definición de la configuración del sistema
type LayerConfig struct {
	Layers Layers `yaml:"layers"`
}

// Definición de las capas del sistema
type Layers struct {
	Domain         []string `yaml:"domain"`
	Application    []string `yaml:"application"`
	Infrastructure []string `yaml:"infrastructure"`
}

// Estructura para almacenar información de los paquetes importados y su capa
type FileImport struct {
	Name  string
	Path  string
	Layer string
}

// Estructura para almacenar información de los archivos y violaciones
type DependencyAnalyzer struct {
	PackagesInfo         []FileImport
	DependencyViolations []string
}

// Inicializar un nuevo analizador de dependencias
func NewDependencyAnalyzer() *DependencyAnalyzer {
	return &DependencyAnalyzer{
		PackagesInfo:         []FileImport{},
		DependencyViolations: []string{},
	}
}

// Añadir información de un archivo importado
func (da *DependencyAnalyzer) AddFileImport(name, path, layer string) {
	fileImport := FileImport{
		Name:  name,
		Path:  path,
		Layer: layer,
	}
	da.PackagesInfo = append(da.PackagesInfo, fileImport)
}

// Añadir una violación de dependencias
func (da *DependencyAnalyzer) AddViolation(violation string) {
	da.DependencyViolations = append(da.DependencyViolations, violation)
}

// Cargar el archivo monitor.yml
func loadLayerConfig(path string) (LayerConfig, error) {
	var config LayerConfig

	content, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Función principal
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <repo_path>")
		return
	}

	repoPath := os.Args[1]

	// Cargar el archivo monitor.yml para las capas
	layerConfig, err := loadLayerConfig(filepath.Join(repoPath, "monitor.yml"))
	if err != nil {
		fmt.Printf("Error loading layer configuration: %v\n", err)
		return
	}

	// Inicializar el analizador de dependencias
	analyzer := NewDependencyAnalyzer()

	// Recorrer todos los archivos y clasificarlos
	err = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			classifyFile(path, layerConfig, analyzer, repoPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", repoPath, err)
		return
	}

	// Realizar el análisis de la inversión de dependencias
	analyzeDependencyInversion(analyzer, layerConfig)

	// Mostrar las violaciones encontradas
	printViolations(analyzer)
}

// Clasificar archivos por capas
func classifyFile(filePath string, config LayerConfig, analyzer *DependencyAnalyzer, repoPath string) {
	absFilePath, _ := filepath.Abs(filePath)

	// Convertimos las rutas relativas de las capas a rutas absolutas
	for _, domainPath := range config.Layers.Domain {
		domainAbsPath, _ := filepath.Abs(filepath.Join(repoPath, domainPath))
		if isSubPath(domainAbsPath, absFilePath) {
			pkgName, _ := getPackageName(filePath)
			analyzer.AddFileImport(pkgName, absFilePath, "domain")
			return
		}
	}

	for _, applicationPath := range config.Layers.Application {
		applicationAbsPath, _ := filepath.Abs(filepath.Join(repoPath, applicationPath))
		if isSubPath(applicationAbsPath, absFilePath) {
			pkgName, _ := getPackageName(filePath)
			analyzer.AddFileImport(pkgName, absFilePath, "application")
			return
		}
	}

	for _, infrastructurePath := range config.Layers.Infrastructure {
		infrastructureAbsPath, _ := filepath.Abs(filepath.Join(repoPath, infrastructurePath))
		if isSubPath(infrastructureAbsPath, absFilePath) {
			pkgName, _ := getPackageName(filePath)
			analyzer.AddFileImport(pkgName, absFilePath, "infrastructure")
			return
		}
	}

	fmt.Printf("File: %s does not belong to any known layer\n", filePath)
}

// Función para verificar si un directorio es subdirectorio de otro
func isSubPath(basePath, targetPath string) bool {
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(relPath, "..") && relPath != "."
}

// Función para obtener el nombre del paquete
func getPackageName(filePath string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", fmt.Errorf("error parsing file: %v", err)
	}
	if node.Name != nil {
		return node.Name.Name, nil
	}
	return "", fmt.Errorf("package name not found in file: %s", filePath)
}

// Función para analizar la inversión de dependencias
func analyzeDependencyInversion(analyzer *DependencyAnalyzer, config LayerConfig) {
	filteredPackages := filterPackagesByLayer(analyzer.PackagesInfo, []string{"domain", "application"})

	for _, file := range filteredPackages {
		// Analizar las importaciones
		// imports, err := getFileImports(file.Path)
		// if err != nil {
		// 	fmt.Printf("Error getting imports for file %s: %v\n", file.Path, err)
		// 	continue
		// }

		// _ = imports

		if file.Layer == "application" {
			// si se declara una variable que no sea de domain, application o lib estandar, esta variable debe ser si o si interaface

			// Listar variables declaradas en el archivo
			//readGoFileLineByLine(file.Path)
			//getVariablesAndTypes(file.Path)
			//listStructFields(file.Path)
			interfaceChecker(file.Path)
		}

		// // Validar las importaciones
		// for _, imp := range imports {
		// 	validateImport(analyzer, file, imp, config)
		// }

		// Analizar los structs para verificar si están usando interfaces o structs concretos
		// analyzeStructs(file.Path, config, analyzer)
	} //listar variables declaradas en el archivo

}

func interfaceChecker(filePath string) error {
	fset := token.NewFileSet()

	// Parsear el archivo
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	// Cargar el paquete de información de tipos
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedTypesInfo,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, filePath)
	if err != nil {
		return fmt.Errorf("error loading package: %w", err)
	}

	if len(pkgs) == 0 {
		return fmt.Errorf("no se encontraron paquetes")
	}

	pkg := pkgs[0]

	// Función para verificar si un tipo es una interfaz
	isInterface := func(expr ast.Expr) bool {
		tv := pkg.TypesInfo.Types[expr]
		if tv.Type == nil {
			return false
		}
		_, ok := tv.Type.Underlying().(*types.Interface)
		return ok
	}

	// Recorrer el AST y analizar las variables
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			// Para declaraciones de variables globales y locales
			for _, name := range x.Names {
				fmt.Printf("Variable: %s, Tipo: ", name.Name)
				if isInterface(x.Type) {
					fmt.Println("es una interfaz")
				} else {
					fmt.Println("no es una interfaz")
				}
			}
		case *ast.AssignStmt:
			// Para asignaciones (variables locales)
			for _, lhs := range x.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					fmt.Printf("Variable: %s, Tipo: ", ident.Name)
					if isInterface(lhs) {
						fmt.Println("es una interfaz")
					} else {
						fmt.Println("no es una interfaz")
					}
				}
			}
		}
		return true
	})

	return nil
}

func readGoFileLineByLine(filePath string) error {
	// Abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Crear un escáner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	// Leer línea por línea
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Línea %d: %s\n", lineNumber, line)
		lineNumber++
	}

	// Verificar si hubo errores al escanear
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func listStructFields(filePath string) error {
	fset := token.NewFileSet()

	// Parsear el archivo
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Verificar si el nodo es una declaración de función
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			fmt.Printf("Función encontrada: %s\n", funcDecl.Name.Name)

			// Obtener los parámetros de la función
			if funcDecl.Type.Params != nil {
				for _, param := range funcDecl.Type.Params.List {
					// Obtener los nombres de los parámetros (pueden ser múltiples)
					for _, name := range param.Names {
						fmt.Printf("  Parámetro: %s, Tipo: ", name.Name)

						// Comprobar el tipo de parámetro (ast.Expr)
						switch expr := param.Type.(type) {
						case *ast.Ident:
							fmt.Printf("%s\n", expr.Name)
						case *ast.SelectorExpr:
							if pkgIdent, ok := expr.X.(*ast.Ident); ok {
								fmt.Printf("%s.%s\n", pkgIdent.Name, expr.Sel.Name)
							}
						case *ast.StarExpr:
							if selector, ok := expr.X.(*ast.SelectorExpr); ok {
								if pkgIdent, ok := selector.X.(*ast.Ident); ok {
									fmt.Printf("*%s.%s\n", pkgIdent.Name, selector.Sel.Name)
								}
							}
						default:
							fmt.Printf("tipo desconocido\n")
						}
					}
				}
			}

			// Recorrer las declaraciones de variables dentro del cuerpo de la función
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				if decl, ok := n.(*ast.DeclStmt); ok {
					if genDecl, ok := decl.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
						for _, spec := range genDecl.Specs {
							if valueSpec, ok := spec.(*ast.ValueSpec); ok {
								for _, name := range valueSpec.Names {
									fmt.Printf("  Variable local encontrada: %s, Tipo: ", name.Name)

									switch expr := valueSpec.Type.(type) {
									case *ast.Ident:
										fmt.Printf("%s\n", expr.Name)
									case *ast.SelectorExpr:
										if pkgIdent, ok := expr.X.(*ast.Ident); ok {
											fmt.Printf("%s.%s\n", pkgIdent.Name, expr.Sel.Name)
										}
									default:
										fmt.Printf("tipo desconocido\n")
									}
								}
							}
						}
					}
				}
				if assign, ok := n.(*ast.AssignStmt); ok {
					for i, lh := range assign.Lhs {
						if ident, ok := lh.(*ast.Ident); ok {
							if i < len(assign.Rhs) {
								switch rhs := assign.Rhs[i].(type) {
								case *ast.Ident:
									fmt.Printf("  Variable asignada: %s, Tipo: %s\n", ident.Name, rhs.Name)
								case *ast.SelectorExpr:
									if pkgIdent, ok := rhs.X.(*ast.Ident); ok {
										fmt.Printf("  Variable asignada: %s, Tipo: %s.%s\n", ident.Name, pkgIdent.Name, rhs.Sel.Name)
									}
								default:
									fmt.Printf("  Variable asignada: %s, Tipo: tipo desconocido\n", ident.Name)
								}
							}
						}
					}
				}
				return true
			})
		}

		// Verificar si el nodo es una declaración general (como variables globales)
		if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valueSpec.Names {
						fmt.Printf("Variable global encontrada: %s, Tipo: ", name.Name)

						switch expr := valueSpec.Type.(type) {
						case *ast.Ident:
							fmt.Printf("%s\n", expr.Name)
						case *ast.SelectorExpr:
							if pkgIdent, ok := expr.X.(*ast.Ident); ok {
								fmt.Printf("%s.%s\n", pkgIdent.Name, expr.Sel.Name)
							}
						default:
							fmt.Printf("tipo desconocido o no especificado\n")
						}
					}
				}
			}
		}

		return true
	})

	return nil
}

// Función para extraer las variables y sus tipos de un archivo Go
func getVariablesAndTypes(filePath string) error {
	// Crear el conjunto de archivos para el parser
	fset := token.NewFileSet()

	// Parsear el archivo
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	// Recorrer el AST y extraer variables y tipos
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			// Para declaraciones de variables
			for _, name := range x.Names {
				fmt.Printf("Variable: %s, Tipo: %s\n", name.Name, fmt.Sprintf("%T", x.Type))
			}
		case *ast.AssignStmt:
			// Para asignaciones (variables dentro de funciones)
			for _, lhs := range x.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					fmt.Printf("Variable: %s, Tipo: Asignación desconocida\n", ident.Name)
				}
			}
		}
		return true
	})

	return nil
}

// Función que filtra los archivos según las capas dadas
func filterPackagesByLayer(packages []FileImport, allowedLayers []string) []FileImport {
	var filtered []FileImport
	for _, pkg := range packages {
		for _, layer := range allowedLayers {
			if pkg.Layer == layer {
				filtered = append(filtered, pkg)
				break // No es necesario continuar, ya que ya encontramos una coincidencia
			}
		}
	}
	return filtered
}

// Función para validar las importaciones
func validateImport(analyzer *DependencyAnalyzer, file FileImport, importPath string, config LayerConfig) {
	switch file.Layer {
	case "domain":
		// Los archivos del dominio solo pueden importar de la librería estándar o de la capa de dominio
		for _, domain := range config.Layers.Domain {
			if !isStdLib(importPath) && !strings.Contains(importPath, domain) {
				analyzer.AddViolation(fmt.Sprintf("Violation in %s: imports %s which is not allowed in domain layer", file.Path, importPath))
			}
		}
	case "application":
		// Los archivos de la aplicación pueden importar de la librería estándar, dominio o aplicación
		for _, domain := range config.Layers.Domain {
			if !isStdLib(importPath) && !strings.Contains(importPath, domain) && !strings.Contains(importPath, config.Layers.Application[0]) {
				// Si importa de infraestructura, debe ser una interfaz
				for _, infrastructure := range config.Layers.Infrastructure {
					if strings.Contains(importPath, infrastructure) && !isInterface(importPath) {
						analyzer.AddViolation(fmt.Sprintf("Violation in %s: imports %s from infrastructure without using an interface", file.Path, importPath))
					}
				}
			}
		}
	}
}

// analyzeStructs analiza los structs en un archivo Go y verifica si usan tipos concretos en lugar de interfaces.
func analyzeStructs(filePath string, config LayerConfig, analyzer *DependencyAnalyzer) {
	// Crear el conjunto de archivos y parsear el archivo Go
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", filePath, err)
		return
	}

	// Recorrer el AST del archivo Go
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			// Verificar si el nodo es un struct
			if structType, ok := x.Type.(*ast.StructType); ok {
				// Recorrer los campos del struct
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						// Imprimir el nombre de la variable que se está analizando
						fmt.Printf("Analizando variable: %s en struct %s\n", fieldName.Name, x.Name.Name)

						// Verificar el tipo del campo
						if ident, ok := field.Type.(*ast.Ident); ok {
							// Si el tipo del campo no es una interfaz y es de una capa no permitida
							if !isInterface(ident.Name) && isConcreteTypeNotAllowed(ident.Name, config) {
								analyzer.AddViolation(fmt.Sprintf("Violation in %s: struct %s uses concrete type %s, which is not allowed", filePath, x.Name.Name, ident.Name))
							}
						}
					}
				}
			}
		}
		return true
	})
}

// isInterface es una función auxiliar que verifica si un tipo es una interfaz (basado en convenciones de nombre)
func isInterface(typeName string) bool {
	// Aquí puedes mejorar esta lógica para verificar si el tipo es una interfaz
	return strings.HasPrefix(typeName, "I") // Convención simple: las interfaces comienzan con "I"
}

// isConcreteTypeNotAllowed verifica si un tipo concreto no está permitido
func isConcreteTypeNotAllowed(typeName string, config LayerConfig) bool {
	// Si el tipo pertenece a "domain" o "application", está permitido
	// Si es de "infrastructure" y no es interfaz, no está permitido
	for _, infrastructure := range config.Layers.Infrastructure {
		if strings.Contains(typeName, infrastructure) {
			return true
		}
	}
	return false
}

// Función para determinar si un paquete es de la librería estándar
func isStdLib(importPath string) bool {
	return !strings.Contains(importPath, ".")
}

// Función para obtener las importaciones de un archivo
func getFileImports(filePath string) ([]string, error) {
	var imports []string
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ImportsOnly)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %v", err)
	}

	for _, imp := range node.Imports {
		imports = append(imports, strings.Trim(imp.Path.Value, "\""))
	}

	return imports, nil
}

// Función para imprimir las violaciones de dependencia encontradas
func printViolations(analyzer *DependencyAnalyzer) {
	if len(analyzer.DependencyViolations) == 0 {
		fmt.Println("No dependency violations found.")
	} else {
		fmt.Println("Dependency violations found:")
		for _, violation := range analyzer.DependencyViolations {
			fmt.Println(violation)
		}
	}
}
