package core

type CalculatorUseCases interface {
	Addition(int32, int32) (int32, error)
	Subtraction(int32, int32) (int32, error)
	Multiplication(int32, int32) (int32, error)
	Division(int32, int32) (int32, error)
}

type calculatorUseCases struct {
}

// instance of calculatorUseCases struct, instances mostly implemented with pointer type and returns
func NewCalculatorUseCases() CalculatorUseCases {
	return &calculatorUseCases{}
}

// implementation of interface method Addition
func (arith calculatorUseCases) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

// implementation of interface method Subtraction
func (arith calculatorUseCases) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

// implementation of interface method Multiplication
func (arith calculatorUseCases) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

// implementation of interface method Division
func (arith calculatorUseCases) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
