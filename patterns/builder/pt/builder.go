package main

import "fmt"

// Person representa o "produto" final no padrão Builder.
// Inclui atributos relacionados tanto com o endereço quanto com o trabalho de uma pessoa.
type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

// PersonBuilder é o construtor principal que inicia o processo de construção do objeto Person.
type PersonBuilder struct {
	person *Person // Uma referência à instância de Person que está sendo construída.
}

// NewPersonBuilder inicializa e retorna uma nova instância de PersonBuilder.
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Build finaliza a construção e retorna o objeto Person construído.
func (it *PersonBuilder) Build() *Person {
	return it.person
}

// Works retorna um construtor específico para configurar os atributos de trabalho de Person.
func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

// Lives retorna um construtor específico para configurar o endereço de Person.
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}

// PersonJobBuilder foca na construção da parte de trabalho da pessoa.
type PersonJobBuilder struct {
	PersonBuilder // Inclui PersonBuilder para acesso ao objeto Person.
}

// Métodos encadeados para configurar o trabalho, permitindo uma interface fluida.
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

// PersonAddressBuilder foca na construção da parte de endereço da pessoa.
type PersonAddressBuilder struct {
	PersonBuilder // Inclui PersonBuilder para acesso ao objeto Person.
}

// Métodos encadeados para configurar o endereço, permitindo uma interface fluida.
func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

// No main, demonstra-se como usar o padrão Builder para configurar e construir um objeto Person.
func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(*person)
}
