package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Function to inspect and list all variables, struct fields, and function parameters in the AST of a Go file
func listVariablesAndStructsAndParams(filePath string) error {
	// Create a new token.FileSet to keep track of position information
	fset := token.NewFileSet()

	// Load the package containing the file to resolve types
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  filepath.Dir(filePath),
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return fmt.Errorf("failed to load package: %w", err)
	}

	if len(pkgs) == 0 {
		return fmt.Errorf("no packages found")
	}

	// Find the correct file in the package
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
		return fmt.Errorf("file not found in package")
	}

	// Map to store imports and their aliases
	imports := make(map[string]string)

	// First, collect all imports
	for _, i := range file.Imports {
		importPath := strings.Trim(i.Path.Value, "\"")
		alias := ""
		if i.Name != nil {
			alias = i.Name.Name
		} else {
			// If there's no alias, use the last part of the import path
			parts := strings.Split(importPath, "/")
			alias = parts[len(parts)-1]
		}
		imports[alias] = importPath
	}

	// Traverse the AST and find variable declarations, struct types, assignments, and function parameters
	ast.Inspect(file, func(n ast.Node) bool {
		// Look for variable declarations (var keyword)
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			for _, spec := range decl.Specs {
				vspec := spec.(*ast.ValueSpec)
				for _, name := range vspec.Names {
					var varType string

					// Handle explicit and inferred types
					if vspec.Type != nil {
						varType = getTypeFromAST(vspec.Type, imports)
					} else {
						obj := pkg.TypesInfo.ObjectOf(name)
						varType = obj.Type().String() // Handle inferred types
					}

					fmt.Printf("Variable declared: %s, Type: %s (line: %d)\n", name.Name, varType, fset.Position(name.Pos()).Line)
				}
			}
		}

		// Look for short variable declarations (:=)
		if assign, ok := n.(*ast.AssignStmt); ok && assign.Tok == token.DEFINE {
			for _, lhs := range assign.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					// Resolve the type using the types package
					obj := pkg.TypesInfo.ObjectOf(ident)
					varType := obj.Type().String()
					fmt.Printf("Variable assigned with :=: %s, Type: %s (line: %d)\n", ident.Name, varType, fset.Position(ident.Pos()).Line)
				}
			}
		}

		// Look for struct type declarations
		if typeDecl, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := typeDecl.Type.(*ast.StructType); ok {
				fmt.Printf("\nStruct found: %s (line: %d)\n", typeDecl.Name.Name, fset.Position(typeDecl.Pos()).Line)
				// Inspect the fields of the struct
				for _, field := range structType.Fields.List {
					var fieldType string
					if len(field.Names) > 0 {
						// Field has a name (i.e., non-embedded field)
						fieldName := field.Names[0].Name
						fieldType = getTypeFromAST(field.Type, imports)

						fmt.Printf("Field: %s, Type: %s\n", fieldName, fieldType)
					} else {
						// Embedded field (anonymous field)
						fieldType = getTypeFromAST(field.Type, imports)
						fmt.Printf("Embedded Field: %s\n", fieldType)
					}
				}
			}
		}

		// Look for function declarations and parameters
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			fmt.Printf("\nFunction found: %s (line: %d)\n", funcDecl.Name.Name, fset.Position(funcDecl.Pos()).Line)
			if funcDecl.Type.Params != nil {
				fmt.Println("Function Parameters:")
				for _, param := range funcDecl.Type.Params.List {
					paramType := getTypeFromAST(param.Type, imports)
					for _, paramName := range param.Names {
						fmt.Printf("  Param: %s, Type: %s\n", paramName.Name, paramType)
					}
				}
			}
		}
		return true
	})

	return nil
}

// Helper function to extract the type of a variable or function parameter from an AST node, including handling pointers and package aliases
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
		// Handle pointer types (e.g., *Type)
		return "*" + getTypeFromAST(t.X, imports)
	case *ast.ArrayType:
		// Handle array types (e.g., []Type)
		return "[]" + getTypeFromAST(t.Elt, imports)
	}
	return "unknown"
}

func main() {
	// File path to the Go file you want to analyze
	filePath := "/home/pablo/Projects/Pablo/github.com/devpablocristo/meli/monitor-projects/regular/internal/core/services/user_service.go"

	// Call the function to list variables, structs, and function parameters in the provided Go file
	err := listVariablesAndStructsAndParams(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
