package main

import "fmt"

type Persona struct {
	nombre   string
	apellido string
	edad     int
	genero   string
	hobbies []struct{
		tipo string
	}
}

type Usuario struct {
	nom_usuario, contraseña string
	datos_personales        Persona
}

func (u *Usuario) saludar() {
	fmt.Printf("¡Hola %s", u.nom_usuario)
}

func (this Usuario) saludar2() {
	fmt.Printf("¡Hola %s", this.nom_usuario)
}

func main() {
	u := Usuario{
		nom_usuario: "juanito",
		contraseña:  "lalala",
		datos_personales: Persona{
			nombre:   "Juan",
			apellido: "Perez",
			edad:     40,
			genero:   "m",
			hobbies: []struct{
					tipo string
					}{
						{"Hello"},
					},
				},
			},
		},

		a := SettlementReportAPIResponse{
			Columns: []struct {
			   Key string `json:"key"`
			}{{"a"}},
		   }

	fmt.Printf("%s", u.datos_personales.nombre)
	u.saludar2()
}
