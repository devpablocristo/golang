/*
Escribe un programa en Go que calcule la cantidad de saltos necesarios para recorrer un conjunto de nubes, representadas como un array de 0's y 1's, donde 0 indica una nube segura y 1 indica una nube peligrosa que no se puede pisar. El jugador puede saltar de una nube segura a otra nube segura que esté a una distancia de una o dos nubes. Imprimr por pantalla el resultado de llamar a la función "saltarEnNubes" con el array "c".
*/
package main

import "fmt"

func main() {
	c := []int{0, 0, 0, 0, 1, 0}

	fmt.Println(saltarEnNubes(c))
}

func saltarEnNubes(c []int) int {

	//fmt.Println(c)

	saltos := 0
	for i := 0; i < len(c); i++ {
		j := i + 2
		if c[j] == 0 {
			saltos++
			i = j
		} else {
			j := i + 1
			if c[j] == 0 {
				saltos++
				i = j
			}
		}
	}

	return saltos
}
