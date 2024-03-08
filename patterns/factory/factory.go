package main

import "fmt"

// Employee es la estructura base que se quiere instanciar a través de la fábrica.
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Enfoque Funcional
// NewEmployeeFactory es una función que retorna otra función.
// Esta función interna crea y retorna un nuevo objeto Employee con los parámetros predefinidos.
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Enfoque Estructural
// EmployeeFactory es una estructura que actúa como una fábrica para crear objetos Employee.
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// NewEmployeeFactory2 inicializa y retorna una nueva instancia de EmployeeFactory.
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// Create es un método de EmployeeFactory que crea y retorna un nuevo objeto Employee
// utilizando los datos almacenados en la instancia de la fábrica.
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	// Creando fábricas para diferentes roles utilizando el enfoque funcional.
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	// Creando objetos Employee a través de las fábricas.
	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	// Creando una fábrica para el rol de CEO utilizando el enfoque estructural.
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// Modificación post-creación de la fábrica, mostrando flexibilidad.
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
