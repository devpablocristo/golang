package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
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
	file, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// Configurar el paquete de información de tipos utilizando `go/packages`
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedDeps,
		Fset: fset,
		Dir:  ".",                                  // Directorio de trabajo
		Env:  append([]string{}, "GO111MODULE=on"), // Asegurarse de que se use módulos de Go
	}

	// Cargar el paquete local (usa el nombre del módulo o ruta relativa)
	pkgs, err := packages.Load(cfg, "regular/internal/...")
	if err != nil {
		log.Fatal(err)
	}

	if len(pkgs) == 0 {
		log.Fatal("No se encontraron paquetes")
	}

	pkg := pkgs[0]

	// Mostrar los errores, si los hay, para verificar problemas de carga de dependencias
	if len(pkg.Errors) > 0 {
		for _, err := range pkg.Errors {
			fmt.Printf("Error cargando el paquete: %v\n", err)
		}
	}

	// Función para verificar si un tipo es una interfaz
	isInterface := func(expr ast.Expr) bool {
		tv := pkg.TypesInfo.Types[expr]
		if tv.Type == nil {
			fmt.Printf("Tipo no encontrado para la expresión: %#v\n", expr)
			return false
		}
		_, ok := tv.Type.Underlying().(*types.Interface)
		return ok
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

						// Comprobar si es una interfaz
						if isInterface(param.Type) {
							fmt.Println(" (interfaz)")
						} else {
							fmt.Println(" (no interfaz)")
						}
					}
				}
			}
		}

		return true
	})
}
