package sdkast

import "go/token"

// Estructuras para almacenar información detallada.
type ParameterInfo struct {
	Name string
	Type string
}

type MethodInfo struct {
	Name         string
	Receiver     string
	InputParams  []ParameterInfo
	OutputParams []ParameterInfo
}

type FunctionInfo struct {
	Name         string
	InputParams  []ParameterInfo
	OutputParams []ParameterInfo
}

// VariableInfo contiene información detallada sobre una variable.
type VariableInfo struct {
	Name     string
	Type     string
	Position token.Position
	IsGlobal bool
	Kind     string
}
