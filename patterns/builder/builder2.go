package main

import "fmt"

// Producto
type Car struct {
	make  string
	model string
}

// Builder
type CarBuilder interface {
	SetMake(make string) CarBuilder
	SetModel(model string) CarBuilder
	Build() Car
}

// Concrete Builder
type carBuilder struct {
	car Car
}

func (b *carBuilder) SetMake(make string) CarBuilder {
	b.car.make = make
	return b
}

func (b *carBuilder) SetModel(model string) CarBuilder {
	b.car.model = model
	return b
}

func (b *carBuilder) Build() Car {
	return b.car
}

// Función de utilidad para crear un nuevo builder
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func main() {
	builder := NewCarBuilder()
	car := builder.SetMake("Toyota").SetModel("Corolla").Build()
	fmt.Println(car)
}
