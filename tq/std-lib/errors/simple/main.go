package main

import (
	"errors"
	"fmt"
)

// main es la función principal del programa.
func main() {
	// Intentamos dividir 10 entre 0, lo cual generará un error.
	resultado, err := dividir(10, 0)
	if err != nil {
		// Si hay un error, lo imprimimos y terminamos la ejecución.
		fmt.Println("Error:", err)
		return
	}

	// Si no hay error, imprimimos el resultado de la división.
	fmt.Println("Resultado de la división:", resultado)
}

// dividir realiza una división y devuelve un error si el divisor es cero.
func dividir(dividendo, divisor int) (int, error) {
	// Verificamos si el divisor es cero.
	if divisor == 0 {
		// Devolvemos cero y un error indicando que no se puede dividir por cero.
		return 0, errors.New("no se puede dividir por cero")
	}

	// Realizamos la división y devolvemos el resultado sin error.
	return dividendo / divisor, nil
}
