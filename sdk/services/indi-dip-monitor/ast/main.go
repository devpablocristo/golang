package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Estructura para almacenar información por archivo
type FileInfo struct {
	Interfaces    []string
	NonInterfaces []string
}

// Función para analizar un paquete completo en lugar de solo un archivo
func analyzeGoPackage(pkgPath string, fileInfo *FileInfo) error {
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		Dir:   pkgPath,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg)
	if err != nil {
		return fmt.Errorf("error al cargar el paquete: %w", err)
	}

	if len(pkgs) == 0 {
		return fmt.Errorf("no se encontraron paquetes en %s", pkgPath)
	}

	pkg := pkgs[0]

	// Recorrer los archivos en el paquete
	for _, syntax := range pkg.Syntax {
		ast.Inspect(syntax, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.ValueSpec:
				// Para cada variable, verificamos si es una interfaz
				for _, name := range x.Names {
					if x.Type != nil && isInterfaceType(pkg.TypesInfo.TypeOf(x.Type)) {
						fileInfo.Interfaces = append(fileInfo.Interfaces, name.Name)
					} else {
						fileInfo.NonInterfaces = append(fileInfo.NonInterfaces, name.Name)
					}
				}
			case *ast.Field: // Para parámetros de funciones o campos de estructuras
				for _, name := range x.Names {
					if x.Type != nil && isInterfaceType(pkg.TypesInfo.TypeOf(x.Type)) {
						fileInfo.Interfaces = append(fileInfo.Interfaces, name.Name)
					} else {
						fileInfo.NonInterfaces = append(fileInfo.NonInterfaces, name.Name)
					}
				}
			}
			return true
		})
	}

	return nil
}

// Función para verificar si un tipo es una interfaz
func isInterfaceType(typ types.Type) bool {
	if typ == nil {
		return false
	}
	_, ok := typ.Underlying().(*types.Interface)
	return ok
}

// Función para imprimir los resultados del análisis
func printResults(fileInfo *FileInfo, filePath string) {
	fmt.Printf("Archivo: %s\n", filePath)

	// Imprimir variables que son interfaces
	fmt.Println("    Interfaces encontradas:")
	for _, iface := range fileInfo.Interfaces {
		fmt.Printf("      - %s\n", iface)
	}

	// Imprimir variables que no son interfaces
	fmt.Println("    No son interfaces:")
	for _, nonIface := range fileInfo.NonInterfaces {
		fmt.Printf("      - %s\n", nonIface)
	}
	fmt.Println("-----------------------------")
}

// Función para buscar el directorio que contiene user_service.go
func findUserServiceGoFile(repoPath string) (string, error) {
	var userServiceDir string
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "user_service.go") {
			userServiceDir = filepath.Dir(path)
			return filepath.SkipDir // Detener la búsqueda cuando encontramos el archivo
		}
		return nil
	})
	if userServiceDir == "" {
		return "", fmt.Errorf("no se encontró el archivo user_service.go")
	}
	return userServiceDir, err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <ruta_del_repositorio>")
		return
	}

	repoPath := os.Args[1]

	// Buscar el directorio del archivo user_service.go
	pkgPath, err := findUserServiceGoFile(repoPath)
	if err != nil {
		fmt.Printf("Error al buscar el archivo: %v\n", err)
		return
	}

	// Almacenar información del archivo user_service.go
	fileInfo := &FileInfo{}

	// Analizar el paquete que contiene user_service.go
	fmt.Println("Analizando el paquete que contiene user_service.go...")
	err = analyzeGoPackage(pkgPath, fileInfo)
	if err != nil {
		fmt.Printf("Error al analizar el paquete: %v\n", err)
		return
	}

	// Imprimir los resultados
	printResults(fileInfo, pkgPath)
}
