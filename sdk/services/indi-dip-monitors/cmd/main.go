package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Definición de la estructura del archivo monitor.yml
type LayerConfig struct {
	Layers struct {
		Domain         string `yaml:"domain"`
		Application    string `yaml:"application"`
		Infrastructure string `yaml:"infrastructure"`
	} `yaml:"layers"`
}

// Estructura para almacenar información de los paquetes importados y su capa
type fileImports struct {
	name  string
	path  string
	layer string
}

var packagesInfo []fileImports
var dependencyViolations []string

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

	// Recorrer todos los archivos y clasificarlos
	err = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			classifyFile(path, layerConfig, repoPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", repoPath, err)
		return
	}

	// Realizar el análisis de la inversión de dependencias
	analyzeDependencyInversion(packagesInfo, layerConfig)

	// Mostrar las violaciones encontradas
	printViolations()
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

func classifyFile(filePath string, config LayerConfig, repoPath string) {
	absFilePath, _ := filepath.Abs(filePath)

	// Convertimos las rutas relativas de las capas a rutas absolutas
	domainPath, _ := filepath.Abs(filepath.Join(repoPath, config.Layers.Domain))
	applicationPath, _ := filepath.Abs(filepath.Join(repoPath, config.Layers.Application))
	infrastructurePath, _ := filepath.Abs(filepath.Join(repoPath, config.Layers.Infrastructure))

	// Clasificamos el archivo según la capa a la que pertenezca
	switch {
	case isSubPath(domainPath, absFilePath):
		pkgName, _ := getPackageName(filePath)
		addFileImport(pkgName, absFilePath, "domain")
	case isSubPath(applicationPath, absFilePath):
		pkgName, _ := getPackageName(filePath)
		addFileImport(pkgName, absFilePath, "application")
	case isSubPath(infrastructurePath, absFilePath):
		pkgName, _ := getPackageName(filePath)
		addFileImport(pkgName, absFilePath, "infrastructure")
	default:
		fmt.Printf("File: %s does not belong to any known layer\n", filePath)
	}
}

func isSubPath(basePath, targetPath string) bool {
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(relPath, "..")
}

func getPackageName(filePath string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", fmt.Errorf("error parsing file: %v", err)
	}
	if node.Name != nil {
		return node.Name.Name, nil
	}
	return "", fmt.Errorf("package name not found")
}

func addFileImport(name, path, layer string) {
	newImport := fileImports{
		name:  name,
		path:  path,
		layer: layer,
	}
	packagesInfo = append(packagesInfo, newImport)
}

// Función para analizar la inversión de dependencias
func analyzeDependencyInversion(packages []fileImports, config LayerConfig) {
	for _, file := range packages {
		imports, err := getFileImports(file.path)
		if err != nil {
			fmt.Printf("Error getting imports for file %s: %v\n", file.path, err)
			continue
		}

		for _, imp := range imports {
			validateImport(file, imp, config)
		}
	}
}

// Función para validar cada importación según las reglas de inversión de dependencias
func validateImport(file fileImports, importPath string, config LayerConfig) {
	switch file.layer {
	case "domain":
		// Los archivos del dominio solo pueden importar de la librería estándar o de la capa de dominio
		if !isStdLib(importPath) && !strings.Contains(importPath, config.Layers.Domain) {
			dependencyViolations = append(dependencyViolations, fmt.Sprintf("Violation in %s: imports %s which is not allowed in domain layer", file.path, importPath))
		}
	case "application":
		// Los archivos de la aplicación pueden importar de la librería estándar, dominio o aplicación
		if !isStdLib(importPath) && !strings.Contains(importPath, config.Layers.Domain) && !strings.Contains(importPath, config.Layers.Application) {
			// Si importa de infraestructura, debe ser una interfaz
			if strings.Contains(importPath, config.Layers.Infrastructure) && !isInterfaceUsed(file.path, importPath) {
				dependencyViolations = append(dependencyViolations, fmt.Sprintf("Violation in %s: imports %s from infrastructure without using an interface", file.path, importPath))
			}
		}
	}
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

// Función para determinar si una interfaz se está usando
func isInterfaceUsed(filePath, importPath string) bool {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return false
	}

	interfaces := make(map[string]bool)
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if _, ok := x.Type.(*ast.InterfaceType); ok {
				interfaces[x.Name.Name] = true
			}
		}
		return true
	})

	return len(interfaces) > 0
}

// Función para imprimir las violaciones de dependencia encontradas
func printViolations() {
	if len(dependencyViolations) == 0 {
		fmt.Println("No dependency violations found.")
	} else {
		fmt.Println("Dependency violations found:")
		for _, violation := range dependencyViolations {
			fmt.Println(violation)
		}
	}
}

// FIXME: si hay violacion de depenencias en regular y dice que no esta mierda del orto
