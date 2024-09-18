package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Estructura para almacenar información sobre variables, parámetros, campos de structs, etc.
type EntityInfo struct {
	Name        string
	Type        string
	Position    int
	Category    string
	IsInterface bool // Campo para indicar si el tipo es una interfaz
}

// Función para inspeccionar y listar todas las variables, campos de structs, parámetros de funciones, tipos de retorno e interfaces en el AST de un archivo Go
func listVariablesStructsParamsAndInterfaces(filePath string) ([]EntityInfo, error) {
	// Crear un nuevo token.FileSet para rastrear información de posición
	fset := token.NewFileSet()

	// Cargar el paquete que contiene el archivo para resolver tipos
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  filepath.Dir(filePath),
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	}

	// Encontrar el archivo correcto en el paquete
	var pkg *packages.Package
	var file *ast.File
	for _, p := range pkgs {
		for _, f := range p.Syntax {
			if pkgs[0].Fset.Position(f.Pos()).Filename == filePath {
				pkg = p
				file = f
				break
			}
		}
	}

	if file == nil {
		return nil, fmt.Errorf("file not found in package")
	}

	// Mapas para almacenar importaciones, variables locales y resultados
	imports := make(map[string]string)
	localVariables := make(map[string]bool) // Rastrear variables locales

	// Lista para almacenar todos los resultados
	var results []EntityInfo

	// Recopilar todas las importaciones
	for _, i := range file.Imports {
		importPath := strings.Trim(i.Path.Value, "\"")
		alias := ""
		if i.Name != nil {
			alias = i.Name.Name
		} else {
			// Si no hay alias, usar la última parte de la ruta de importación
			parts := strings.Split(importPath, "/")
			alias = parts[len(parts)-1]
		}
		imports[alias] = importPath
	}

	// Recorrer el AST y encontrar declaraciones de variables, structs, parámetros de funciones, tipos de retorno e interfaces
	ast.Inspect(file, func(n ast.Node) bool {
		// Considerar solo declaraciones de variables a nivel superior como variables globales
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			if _, inFunction := findFunctionScope(n); !inFunction {
				// Estamos fuera de cualquier función o método, por lo que estas son variables globales
				for _, spec := range decl.Specs {
					if vspec, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range vspec.Names {
							// Si esta variable fue declarada localmente, omitirla
							if localVariables[name.Name] {
								continue
							}

							var varType string
							if vspec.Type != nil {
								varType = getTypeFromAST(vspec.Type, imports)
							} else {
								obj := pkg.TypesInfo.ObjectOf(name)
								varType = obj.Type().String() // Manejar tipos inferidos
							}

							isInterface := isInterface(vspec.Type, pkg)

							// Agregar la variable global a los resultados
							results = append(results, EntityInfo{
								Name:        name.Name,
								Type:        varType,
								Position:    fset.Position(name.Pos()).Line,
								Category:    "Global Variable",
								IsInterface: isInterface,
							})
						}
					}
				}
			}
		}

		// Buscar declaraciones de funciones y manejar variables locales dentro del cuerpo de la función
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			// Recopilar parámetros de entrada
			if funcDecl.Type.Params != nil {
				for _, param := range funcDecl.Type.Params.List {
					paramType := getTypeFromAST(param.Type, imports)
					isInterface := isInterface(param.Type, pkg)
					for _, paramName := range param.Names {
						results = append(results, EntityInfo{
							Name:        paramName.Name,
							Type:        paramType,
							Position:    fset.Position(paramName.Pos()).Line,
							Category:    "Function Parameter",
							IsInterface: isInterface,
						})
					}
				}
			}

			// Recopilar tipos de retorno
			if funcDecl.Type.Results != nil {
				for _, result := range funcDecl.Type.Results.List {
					resultType := getTypeFromAST(result.Type, imports)
					isInterface := isInterface(result.Type, pkg)
					results = append(results, EntityInfo{
						Type:        resultType,
						Position:    fset.Position(result.Pos()).Line,
						Category:    "Function Return Type",
						IsInterface: isInterface,
					})
				}
			}

			// Recorrer declaraciones de variables locales dentro del cuerpo de la función
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				if declStmt, ok := n.(*ast.DeclStmt); ok {
					if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
						for _, spec := range genDecl.Specs {
							if vspec, ok := spec.(*ast.ValueSpec); ok {
								for _, name := range vspec.Names {
									var varType string
									if vspec.Type != nil {
										varType = getTypeFromAST(vspec.Type, imports)
									} else {
										obj := pkg.TypesInfo.ObjectOf(name)
										varType = obj.Type().String() // Manejar tipos inferidos
									}

									// Registrar la variable como local
									localVariables[name.Name] = true

									isInterface := isInterface(vspec.Type, pkg)

									// Agregar la variable local a los resultados
									results = append(results, EntityInfo{
										Name:        name.Name,
										Type:        varType,
										Position:    fset.Position(name.Pos()).Line,
										Category:    "Local Variable",
										IsInterface: isInterface,
									})
								}
							}
						}
					}
				}
				// Manejar declaraciones cortas (:=)
				if assign, ok := n.(*ast.AssignStmt); ok && assign.Tok == token.DEFINE {
					for _, lhs := range assign.Lhs {
						if ident, ok := lhs.(*ast.Ident); ok {
							obj := pkg.TypesInfo.ObjectOf(ident)
							varType := obj.Type().String()

							// Registrar la variable como local
							localVariables[ident.Name] = true

							isInterface := isInterface(ident, pkg)

							// Agregar la variable local a los resultados
							results = append(results, EntityInfo{
								Name:        ident.Name,
								Type:        varType,
								Position:    fset.Position(ident.Pos()).Line,
								Category:    "Local Variable",
								IsInterface: isInterface,
							})
						}
					}
				}
				return true
			})
		}

		// Buscar declaraciones de structs
		if typeDecl, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := typeDecl.Type.(*ast.StructType); ok {
				// Agregar los campos del struct a los resultados
				for _, field := range structType.Fields.List {
					var fieldType string
					if len(field.Names) > 0 {
						fieldName := field.Names[0].Name
						fieldType = getTypeFromAST(field.Type, imports)
						isInterface := isInterface(field.Type, pkg)

						results = append(results, EntityInfo{
							Name:        fieldName,
							Type:        fieldType,
							Position:    fset.Position(field.Pos()).Line,
							Category:    "Struct Field",
							IsInterface: isInterface,
						})
					}
				}
			}
		}

		return true
	})

	return results, nil
}

// Función de ayuda para verificar si un tipo dado es una interfaz
func isInterface(expr ast.Expr, pkg *packages.Package) bool {
	if typ, ok := pkg.TypesInfo.Types[expr]; ok {
		_, isInterface := typ.Type.Underlying().(*types.Interface)
		return isInterface
	}
	return false
}

// Función de ayuda para extraer el tipo de una variable o parámetro de función desde un nodo AST, incluyendo manejo de punteros y alias de paquetes
func getTypeFromAST(expr ast.Expr, imports map[string]string) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if pkgIdent, ok := t.X.(*ast.Ident); ok {
			pkgAlias := pkgIdent.Name
			if pkgPath, ok := imports[pkgAlias]; ok {
				return fmt.Sprintf("%s.%s", pkgPath, t.Sel.Name)
			}
			return fmt.Sprintf("%s.%s", pkgAlias, t.Sel.Name)
		}
	case *ast.StarExpr:
		// Manejar tipos punteros (e.g., *Type)
		return "*" + getTypeFromAST(t.X, imports)
	case *ast.ArrayType:
		// Manejar tipos de arrays (e.g., []Type)
		return "[]" + getTypeFromAST(t.Elt, imports)
	}
	return "unknown"
}

// Función de ayuda para determinar si estamos dentro de un alcance de función
func findFunctionScope(n ast.Node) (string, bool) {
	if funcDecl, ok := n.(*ast.FuncDecl); ok {
		return funcDecl.Name.Name, true
	}
	return "", false
}

// Función que verifica si el tipo es primitivo (int, string, etc.)
func isPrimitiveType(varType string) bool {
	primitives := []string{"int", "string", "bool", "float32", "float64", "byte", "rune", "complex64", "complex128"}
	for _, primitive := range primitives {
		if varType == primitive {
			return true
		}
	}
	return false
}

func main() {
	// Ruta al archivo Go que quieres analizar
	filePath := "/home/pablo/Projects/Pablo/github.com/devpablocristo/meli/monitor-projects/regular/internal/core/services/user_service.go"

	// Llamar a la función para listar variables, structs, parámetros de funciones, tipos de retorno e interfaces en el archivo Go proporcionado
	results, err := listVariablesStructsParamsAndInterfaces(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Imprimir solo los resultados que no sean interfaces, no sean tipos primitivos, y cuya categoría sea relevante
	for _, entity := range results {
		if !entity.IsInterface && !isPrimitiveType(entity.Type) && (entity.Category == "Global Variable" || entity.Category == "Local Variable" || entity.Category == "Function Parameter" || entity.Category == "Function Return Type" || entity.Category == "Struct Field") {
			fmt.Printf("Category: %s, Name: %s, Type: %s, Line: %d\n", entity.Category, entity.Name, entity.Type, entity.Position)
		}
	}
}
