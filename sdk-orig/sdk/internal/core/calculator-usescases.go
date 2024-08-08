package core

type CalculatorUseCasesPort interface {
	Addition(int32, int32) (int32, error)
	Subtraction(int32, int32) (int32, error)
	Multiplication(int32, int32) (int32, error)
	Division(int32, int32) (int32, error)
}

type CalculatorUseCases struct {
}

// instance of CalculatorUseCases struct, instances mostly implemented with pointer type and returns
func NewCalculatorUseCases() CalculatorUseCasesPort {
	return &CalculatorUseCases{}
}

// implementation of interface method Addition
func (arith CalculatorUseCases) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

// implementation of interface method Subtraction
func (arith CalculatorUseCases) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

// implementation of interface method Multiplication
func (arith CalculatorUseCases) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

// implementation of interface method Division
func (arith CalculatorUseCases) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
