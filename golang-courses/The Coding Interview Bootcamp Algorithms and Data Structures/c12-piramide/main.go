package main

import "fmt"

func main() {

	piramide(7)
}

// funcinan las 2 formas
func piramide(esc int) {

	// esc por escalones
	// i representa el número de escalones
	for i := 0; i < esc; i++ {
		// j representa el número de espacios antes del primer '#'
		for j := 0; j < esc-i; j++ {
			fmt.Print(" ")
		}
		// k representa el ńumero de '#' a imprimirse
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("#")
		}
		// salto de línea
		fmt.Println()
	}
}

/*func piramide(esc int) {

	// esc por escalones
	// i representa el número de escalones
	for i := 1; i <= esc; i++ {
		// j representa el número de espacios antes del primer '#'
		for j := 1; j <= esc-i; j++ {
			fmt.Print(" ")
		}
		// k representa el ńumero de '#' a imprimirse
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("#")
		}
		// salto de línea
		fmt.Println()
	}
}*/
