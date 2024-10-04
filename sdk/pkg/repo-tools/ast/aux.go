package sdkast

import (
	"fmt"
	"go/ast"
)

// Función auxiliar para convertir ast.Expr a string.
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	case *ast.FuncType:
		return "func"
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.Ellipsis:
		return "..." + exprToString(t.Elt)
	default:
		return fmt.Sprintf("%T", t)
	}
}

// Función auxiliar para obtener información de parámetros.
func getParameterInfo(fl *ast.FieldList) []ParameterInfo {
	var params []ParameterInfo
	if fl == nil {
		return params
	}
	for _, field := range fl.List {
		typeStr := exprToString(field.Type)
		if len(field.Names) == 0 {
			// Parámetro anónimo
			params = append(params, ParameterInfo{
				Name: "",
				Type: typeStr,
			})
		} else {
			for _, name := range field.Names {
				params = append(params, ParameterInfo{
					Name: name.Name,
					Type: typeStr,
				})
			}
		}
	}
	return params
}

// Función auxiliar para recoger nodos AST basados en un filtro.
func collectNodes[T any](node ast.Node, filter func(ast.Node) (T, bool)) ([]T, error) {
	var results []T
	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		if result, ok := filter(n); ok {
			results = append(results, result)
		}
		return true
	})
	return results, nil
}
