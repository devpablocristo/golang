package main

import "fmt"

type animal struct {
	hablar() string
}

type humano struct {
	edad   int
	nombre string
	genero string
}

type perro struct {
	edad int
	nombre string
	raza  string
	sana  bool
	dueño humano
}

type gato struct {
	colorOjos string
	genero string
	raza string
}



func main() {
	humano := {
		nombre: "Luis",
		genero: "Alvarez",
		edad:   50,
	}
}
