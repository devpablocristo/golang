package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// El archivo de ejemplo
	src := `
package services

import (
	apiclt "regular/internal/adapters/driven/api-client"
	"regular/internal/adapters/driven/auth"
	"regular/internal/adapters/driven/db"
	"regular/internal/core/domain"
)

var globalVariable = "This is a global variable"
var anotherVariable int

type UserUseCases interface {
	GetUser(string) (*domain.User, error)
	SaveUser(*domain.User) error
}

type useUserCases struct {
	repo db.Repository
	auth auth.AuthService
	aclt *apiclt.ApiClient
}

func NewUserUseCases(repo db.Repository, auth auth.AuthService, aclt *apiclt.ApiClient) UserUseCases {
	return &useUserCases{
		repo: repo,
		auth: auth,
		aclt: aclt,
	}
}

func (s *useUserCases) GetUser(id string) (*domain.User, error) {
	var okVar *domain.User
	_ = okVar
	return s.repo.GetByID(id)
}

func (s *useUserCases) SaveUser(user *domain.User) error {
	var wrongVar apiclt.ApiClient
	_ = wrongVar
	return s.repo.Save(user)
}
`

	// Crear un conjunto de archivos de posición
	fset := token.NewFileSet()

	// Analizar el archivo fuente
	file, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Recorrer el AST para funciones y variables
	ast.Inspect(file, func(n ast.Node) bool {
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
}
