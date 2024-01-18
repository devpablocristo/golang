package main

import (
	"fmt"
	"math"
)

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	//var numero int

	//fmt.Print("Digite um número: ")
	//fmt.Scan(&numero)

	numero := 10
	//numero := 11

	if ehPrimo(numero) {
		fmt.Printf("%d é um número primo.\n", numero)
	} else {
		fmt.Printf("%d não é um número primo.\n", numero)
	}
}
