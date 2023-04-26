package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello!")
}

func main() {

	p := person{
		name:     "John",
		lastname: "Lennon",
		age:      32,
	}

	// este no funciona
	// porque speck tiene un reciber de tipo puntero
	// entonces p tiene que ser tb de tipo puntero para
	// poder funcionar correctamente
	//saySomething(p)

	saySomething(&p)

	// aqui si funcional, pq es solo el tipo person
	p.speak()

}

func saySomething(h human) {
	h.speak()
}
