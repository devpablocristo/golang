package main

import "fmt"

type Persona struct {
	edad           int
	nombre, genero string
}

type Mascota struct {
	edad         int
	nombre, raza string
	sana         bool
	dueño        Persona
}

/*
Nuevo método
*/
func (m *Mascota) saludar() {
	fmt.Printf("¡Hola %s! Bienvenido a casa\n", m.dueño.nombre)
}

func (m *Mascota) ladrar() {
	fmt.Printf("¡guau! me llamo %s y tengo %d años\n", m.nombre, m.edad)
}

func (m *Mascota) Edad() int {
	return m.edad
}

func (m *Mascota) SetEdad(nuevaEdad int) {
	m.edad = nuevaEdad
}

func main() {
	mascota := Mascota{
		dueño: Persona{
			nombre: "Luis",
			genero: "M",
			edad:   50,
		},
		edad:   10,
		nombre: "Guayaba",
		raza:   "Tampoco la conozco",
		sana:   true,
	}
	mascota.saludar()
	mascota.ladrar()
}
