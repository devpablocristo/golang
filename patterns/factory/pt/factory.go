package main

import "fmt"

// Employee é a estrutura base que queremos instanciar através da fábrica.
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Abordagem Funcional
// NewEmployeeFactory é uma função que retorna outra função.
// Esta função interna cria e retorna um novo objeto Employee com os parâmetros pré-definidos.
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Abordagem Estrutural
// EmployeeFactory é uma estrutura que atua como uma fábrica para criar objetos Employee.
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// NewEmployeeFactory2 inicializa e retorna uma nova instância de EmployeeFactory.
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// Create é um método de EmployeeFactory que cria e retorna um novo objeto Employee
// utilizando os dados armazenados na instância da fábrica.
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	// Criando fábricas para diferentes papéis usando a abordagem funcional.
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	// Criando objetos Employee através das fábricas.
	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	// Criando uma fábrica para o papel de CEO usando a abordagem estrutural.
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// Modificação pós-criação da fábrica, mostrando flexibilidade.
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
