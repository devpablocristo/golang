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

// Estructura para almacenar información por paquete
type PackageInfo struct {
	Interfaces      map[string]bool
	StandardVars    map[string]bool
	NonStandardVars map[string]string // Mapa de variable y su paquete de origen
}

// Función para analizar un paquete completo
func analyzeGoPackage(pkgPath string, pkgInfo *PackageInfo) error {
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
				// Ignorar los valores de las funciones
				for _, name := range x.Names {
					if !isFunctionOrReceiver(name.Name) {
						if x.Type != nil && isInterfaceType(pkg.TypesInfo.TypeOf(x.Type)) {
							pkgInfo.Interfaces[name.Name] = true
						} else {
							categorizeVariable(name.Name, pkg.TypesInfo.TypeOf(x.Type), pkgInfo)
						}
					}
				}
			case *ast.Field: // Para parámetros de funciones o campos de estructuras
				for _, name := range x.Names {
					if !isFunctionOrReceiver(name.Name) {
						if x.Type != nil && isInterfaceType(pkg.TypesInfo.TypeOf(x.Type)) {
							pkgInfo.Interfaces[name.Name] = true
						} else {
							categorizeVariable(name.Name, pkg.TypesInfo.TypeOf(x.Type), pkgInfo)
						}
					}
				}
			}
			return true
		})
	}

	return nil
}

// Función para determinar si un nombre corresponde a una función o un receiver
func isFunctionOrReceiver(name string) bool {
	return strings.HasPrefix(name, "Get") || strings.HasPrefix(name, "Save") || name == "s"
}

// Función para categorizar una variable entre estándar o no estándar
func categorizeVariable(varName string, typ types.Type, pkgInfo *PackageInfo) {
	if typ == nil {
		return
	}

	if isStandardPackage(typ) {
		pkgInfo.StandardVars[varName] = true
	} else {
		pkgPath := getPackagePath(typ)
		pkgInfo.NonStandardVars[varName] = pkgPath
	}
}

// Función para verificar si un tipo es parte de la librería estándar
func isStandardPackage(typ types.Type) bool {
	if typ == nil {
		return false
	}
	pkg := typ.String()
	return !strings.Contains(pkg, ".")
}

// Obtener el paquete de una variable no estándar
func getPackagePath(typ types.Type) string {
	if named, ok := typ.(*types.Named); ok {
		if pkg := named.Obj().Pkg(); pkg != nil {
			return pkg.Path()
		}
	}
	return "desconocido"
}

// Función para verificar si un tipo es una interfaz
func isInterfaceType(typ types.Type) bool {
	if typ == nil {
		return false
	}
	_, ok := typ.Underlying().(*types.Interface)
	return ok
}

// Función para imprimir los resultados del análisis por paquete
func printResults(pkgInfo *PackageInfo, pkgPath string) {
	fmt.Printf("Paquete: %s\n", pkgPath)

	// Imprimir variables que son interfaces
	fmt.Println("    Interfaces encontradas:")
	for iface := range pkgInfo.Interfaces {
		fmt.Printf("      - %s\n", iface)
	}

	// Imprimir variables estándar
	fmt.Println("    Variables estándar:")
	for stdVar := range pkgInfo.StandardVars {
		fmt.Printf("      - %s\n", stdVar)
	}

	// Imprimir variables no estándar con su paquete
	fmt.Println("    Variables no estándar:")
	for varName, pkgPath := range pkgInfo.NonStandardVars {
		fmt.Printf("      - %s (paquete: %s)\n", varName, pkgPath)
	}
	fmt.Println("-----------------------------")
}

// Función para buscar el paquete que contiene el archivo user_service.go
func findUserServicePackage(repoPath string) (string, error) {
	var userServicePkg string
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "user_service.go") {
			userServicePkg = filepath.Dir(path)
			return filepath.SkipDir // Detener la búsqueda cuando encontramos el archivo
		}
		return nil
	})
	if userServicePkg == "" {
		return "", fmt.Errorf("no se encontró el archivo user_service.go")
	}
	return userServicePkg, err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <ruta_del_repositorio>")
		return
	}

	repoPath := os.Args[1]

	// Buscar el paquete del archivo user_service.go
	pkgPath, err := findUserServicePackage(repoPath)
	if err != nil {
		fmt.Printf("Error al buscar el paquete: %v\n", err)
		return
	}

	// Almacenar información del paquete
	pkgInfo := &PackageInfo{
		Interfaces:      make(map[string]bool),
		StandardVars:    make(map[string]bool),
		NonStandardVars: make(map[string]string),
	}

	// Analizar el paquete completo
	fmt.Println("Analizando el paquete que contiene user_service.go...")
	err = analyzeGoPackage(pkgPath, pkgInfo)
	if err != nil {
		fmt.Printf("Error al analizar el paquete: %v\n", err)
		return
	}

	// Imprimir los resultados por paquete
	printResults(pkgInfo, pkgPath)
}
