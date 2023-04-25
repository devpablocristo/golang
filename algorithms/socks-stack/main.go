/*
Crear un programa que reciba una pila de medias representada por un slice de enteros y cuenta cuántos pares hay en la pila. Cada número en la pila representa una media, y un par se forma cuando se encuentran dos medias del mismo número.
La salida del programa será un número entero que indica la cantidad de pares que se encontraron en la pila de medias representada por el slice de enteros.
*/

package main

import "fmt"

func main() {

	// ar es una pila de medias, cada numero es un media, encotrar cuantos pares hay
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(paresEnArray(ar))
}

func paresEnArray(arr []int) int {
	//Create a   dictionary of values for each element
	m := make(map[int]int)

	pares := 0
	for _, num := range arr {
		m[num] += 1

		if m[num]%2 == 0 {
			pares++
			m[num] = 0
		}
	}
	//fmt.Println(m)

	return pares
}
