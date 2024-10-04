package sdkast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
	"sync"

	ports "github.com/devpablocristo/golang/sdk/pkg/repo-tools/ast/ports"
	"golang.org/x/tools/go/packages"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	config ports.Config
}

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			initError = err
			return
		}

		instance = &service{
			config: config,
		}
	})
	return instance, initError
}

// Función auxiliar para parsear un archivo Go y obtener el AST.
func (s *service) parseFile(filePath string, mode parser.Mode) (*ast.File, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	node, err := parser.ParseFile(token.NewFileSet(), filePath, src, mode)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %w", err)
	}
	return node, nil
}

// ReadImports analiza un archivo Go y devuelve los imports encontrados.
func (s *service) ReadImports(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var imports []string
	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		imports = append(imports, importPath)
	}

	return imports, nil
}

// ReadFunctions analiza un archivo Go y devuelve las funciones encontradas.
func (s *service) ReadFunctions(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (string, bool) {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Recv == nil {
			return fn.Name.Name, true
		}
		return "", false
	}

	return collectNodes(node, filter)
}

// ReadMethods analiza un archivo Go y devuelve los métodos encontrados.
func (s *service) ReadMethods(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (string, bool) {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Recv != nil {
			return fn.Name.Name, true
		}
		return "", false
	}

	return collectNodes(node, filter)
}

// ReadStructs analiza un archivo Go y devuelve las structs encontradas.
func (s *service) ReadStructs(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (string, bool) {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if _, isStruct := ts.Type.(*ast.StructType); isStruct {
				return ts.Name.Name, true
			}
		}
		return "", false
	}

	return collectNodes(node, filter)
}

// ReadInterfaces analiza un archivo Go y devuelve las interfaces encontradas.
func (s *service) ReadInterfaces(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (string, bool) {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if _, isInterface := ts.Type.(*ast.InterfaceType); isInterface {
				return ts.Name.Name, true
			}
		}
		return "", false
	}

	return collectNodes(node, filter)
}

// ReadConstants analiza un archivo Go y devuelve las constantes encontradas.
func (s *service) ReadConstants(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) ([]string, bool) {
		if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
			var constants []string
			for _, spec := range genDecl.Specs {
				if valSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valSpec.Names {
						constants = append(constants, name.Name)
					}
				}
			}
			return constants, len(constants) > 0
		}
		return nil, false
	}

	constantLists, err := collectNodes(node, filter)
	if err != nil {
		return nil, err
	}

	// Aplanar la lista de listas
	var constants []string
	for _, list := range constantLists {
		constants = append(constants, list...)
	}

	return constants, nil
}

// ReadVariables analiza un archivo Go y devuelve las variables globales encontradas.
func (s *service) ReadVariables(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) ([]string, bool) {
		if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
			var variables []string
			for _, spec := range genDecl.Specs {
				if valSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valSpec.Names {
						variables = append(variables, name.Name)
					}
				}
			}
			return variables, len(variables) > 0
		}
		return nil, false
	}

	variableLists, err := collectNodes(node, filter)
	if err != nil {
		return nil, err
	}

	// Aplanar la lista de listas
	var variables []string
	for _, list := range variableLists {
		variables = append(variables, list...)
	}

	return variables, nil
}

// ReadTypeAliases analiza un archivo Go y devuelve los type aliases encontrados.
func (s *service) ReadTypeAliases(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (string, bool) {
		if ts, ok := n.(*ast.TypeSpec); ok && ts.Assign != 0 {
			return ts.Name.Name, true
		}
		return "", false
	}

	return collectNodes(node, filter)
}

// ReadPackageName analiza un archivo Go y devuelve el nombre del paquete.
func (s *service) ReadPackageName(filePath string) (string, error) {
	node, err := s.parseFile(filePath, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	return node.Name.Name, nil
}

// CountStatements analiza un archivo Go y cuenta el número de declaraciones (statements).
func (s *service) CountStatements(filePath string) (int, error) {
	node, err := s.parseFile(filePath, parser.ParseComments)
	if err != nil {
		return 0, err
	}

	count := 0
	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(ast.Stmt); ok {
			count++
		}
		return true
	})

	return count, nil
}

// ReadComments analiza un archivo Go y devuelve los comentarios encontrados.
func (s *service) ReadComments(filePath string) ([]string, error) {
	node, err := s.parseFile(filePath, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var comments []string
	for _, commentGroup := range node.Comments {
		comments = append(comments, commentGroup.Text())
	}

	return comments, nil
}

// ReadMethodsInfo analiza un archivo Go y devuelve información detallada de los métodos encontrados.
func (s *service) ReadMethodsInfo(filePath string) ([]MethodInfo, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (MethodInfo, bool) {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Recv != nil {
			methodInfo := MethodInfo{
				Name: fn.Name.Name,
			}
			if len(fn.Recv.List) > 0 {
				methodInfo.Receiver = exprToString(fn.Recv.List[0].Type)
			}
			methodInfo.InputParams = getParameterInfo(fn.Type.Params)
			methodInfo.OutputParams = getParameterInfo(fn.Type.Results)
			return methodInfo, true
		}
		return MethodInfo{}, false
	}

	return collectNodes(node, filter)
}

// ReadFunctionsInfo analiza un archivo Go y devuelve información detallada de las funciones encontradas.
func (s *service) ReadFunctionsInfo(filePath string) ([]FunctionInfo, error) {
	node, err := s.parseFile(filePath, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	filter := func(n ast.Node) (FunctionInfo, bool) {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Recv == nil {
			funcInfo := FunctionInfo{
				Name:         fn.Name.Name,
				InputParams:  getParameterInfo(fn.Type.Params),
				OutputParams: getParameterInfo(fn.Type.Results),
			}
			return funcInfo, true
		}
		return FunctionInfo{}, false
	}

	return collectNodes(node, filter)
}

// ReadVariablesDetailed analiza un archivo Go y devuelve información detallada de las variables.
func (s *service) ReadVariablesDetailed(filePath string) ([]VariableInfo, error) {
	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo, Dir: ""}
	pkgs, err := packages.Load(cfg, fmt.Sprintf("file=%s", filePath))
	if err != nil || len(pkgs) == 0 {
		return nil, fmt.Errorf("error loading package: %w", err)
	}

	var results []VariableInfo
	pkg := pkgs[0]
	file, err := s.getFileFromPackage(pkg, filePath)
	if err != nil {
		return nil, err
	}

	imports := s.extractImports(file)

	ast.Inspect(file, func(n ast.Node) bool {
		switch decl := n.(type) {
		case *ast.GenDecl:
			s.processVariables(decl, pkg, imports, &results)
		case *ast.FuncDecl:
			s.processFunctionAssignments(decl, pkg, &results)
		}
		return true
	})

	return results, nil
}

// processVariables procesa las variables globales en el archivo.
func (s *service) processVariables(decl *ast.GenDecl, pkg *packages.Package, imports map[string]string, results *[]VariableInfo) {
	for _, spec := range decl.Specs {
		if vspec, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range vspec.Names {
				varType := s.getVariableType(vspec.Type, name, pkg, imports)
				isGlobal := true
				kind := s.getKindFromObj(pkg.TypesInfo.ObjectOf(name))

				pos := pkg.Fset.Position(name.Pos())
				*results = append(*results, VariableInfo{
					Name:     name.Name,
					Type:     varType,
					Position: pos,
					IsGlobal: isGlobal,
					Kind:     kind,
				})
			}
		}
	}
}

// processFunctionAssignments procesa las variables locales dentro de funciones.
func (s *service) processFunctionAssignments(funcDecl *ast.FuncDecl, pkg *packages.Package, results *[]VariableInfo) {
	fset := pkg.Fset
	if fset == nil {
		fset = token.NewFileSet()
	}

	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if assignStmt, ok := n.(*ast.AssignStmt); ok {
			for _, lhs := range assignStmt.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					obj := pkg.TypesInfo.ObjectOf(ident)
					if obj != nil {
						kind := s.getKindFromObj(obj)
						varType := obj.Type().String()
						isGlobal := false
						pos := fset.Position(ident.Pos())
						*results = append(*results, VariableInfo{
							Name:     ident.Name,
							Type:     varType,
							Position: pos,
							IsGlobal: isGlobal,
							Kind:     kind,
						})
					}
				}
			}
		}
		return true
	})
}

// getVariableType obtiene el tipo de una variable.
func (s *service) getVariableType(expr ast.Expr, name *ast.Ident, pkg *packages.Package, imports map[string]string) string {
	if expr != nil {
		return s.getTypeFromAST(expr, imports)
	}
	return pkg.TypesInfo.ObjectOf(name).Type().String()
}

// getKindFromObj devuelve el tipo (struct, interface) de un objeto.
func (s *service) getKindFromObj(obj types.Object) string {
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

// getTypeFromAST obtiene el tipo desde una expresión AST.
func (s *service) getTypeFromAST(expr ast.Expr, imports map[string]string) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if pkgIdent, ok := t.X.(*ast.Ident); ok {
			if pkgPath, exists := imports[pkgIdent.Name]; exists {
				return fmt.Sprintf("%s.%s", pkgPath, t.Sel.Name)
			}
			return fmt.Sprintf("%s.%s", pkgIdent.Name, t.Sel.Name)
		}
	case *ast.StarExpr:
		return "*" + s.getTypeFromAST(t.X, imports)
	case *ast.ArrayType:
		return "[]" + s.getTypeFromAST(t.Elt, imports)
	}
	return "unknown"
}

// extractImports extrae las importaciones del archivo.
func (s *service) extractImports(file *ast.File) map[string]string {
	imports := make(map[string]string)
	for _, i := range file.Imports {
		alias := ""
		if i.Name != nil {
			alias = i.Name.Name
		} else {
			path := strings.Trim(i.Path.Value, "\"")
			parts := strings.Split(path, "/")
			alias = parts[len(parts)-1]
		}
		imports[alias] = strings.Trim(i.Path.Value, "\"")
	}
	return imports
}

// getFileFromPackage obtiene el archivo AST desde el paquete.
func (s *service) getFileFromPackage(pkg *packages.Package, filePath string) (*ast.File, error) {
	for _, file := range pkg.Syntax {
		if pkg.Fset.Position(file.Pos()).Filename == filePath {
			return file, nil
		}
	}
	return nil, fmt.Errorf("archivo no encontrado en el paquete")
}
