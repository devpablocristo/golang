// Golang program to illustrate the usage of
// fmt.Scanf() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some variables
	var nombre string
	var apellido string
	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanf("%s", &nombre)
	fmt.Scanf("%d", &apellido)

	// Printing the given texts
	fmt.Printf("El nombre es %s y el apellido es %s.",
		nombre, apellido)
}
