package main

import "fmt"

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
	// Contacto informacionDeContacto // tambien es valido
	informacionDeContacto // solo declarar esto es suficiente
}

type informacionDeContacto struct {
	Telefono     int
	Email        string
	CodigoPostal int
}

func main() {
	/*
		// valido si se declara "Contacto"
		p := persona{
			Nombre:   "Homero",
			Apellido: "Simpson",
			Edad:     39,
			Contacto: informacionDeContacto{
				Telefono:     12345656,
				Email:        "hsimpson@burnsplant.com",
				CodigoPostal: 88888,
			},
		}
	*/

	p := persona{
		Nombre:   "Homero",
		Apellido: "Simpson",
		Edad:     39,
		informacionDeContacto: informacionDeContacto{
			Telefono:     12345656,
			Email:        "hsimpson@burnsplant.com",
			CodigoPostal: 88888,
		},
	}

	fmt.Println(p)
	fmt.Println(p.Nombre, p.Apellido, p.Edad)

	p.Nombre = "Homero J."

	fmt.Println(p)
	fmt.Println(p.Nombre, p.Apellido, p.Edad)
	fmt.Printf("%+v\n", p)

	fmt.Println("--------------------")
	p.obternerNombre()

	punteroP := &p
	punteroP.actualizarNombre("Marge")
	fmt.Println("--------------------")
	p.obternerNombre()

	// tb es válido
	// y requiere menos código
	p.actualizarNombre("Lisa")
	fmt.Println("--------------------")
	p.obternerNombre()

	miSlice := []string{"Hola", ",", "como", "estas", "?"}

	actualizarSlice(miSlice)

	fmt.Println("--------------------")
	fmt.Println(miSlice)

}

// como se puede ver no hizo falta hacer nada relativo a punteros
// para que se apliue en la variable de origen
// los slices son punteros
func actualizarSlice(ss []string) {
	ss[0] = "Hello"
}

func (punteroP *persona) actualizarNombre(nuevoNombre string) {
	(*punteroP).Nombre = nuevoNombre
}

func (p persona) obternerNombre() {
	fmt.Println(p.Nombre)
}
