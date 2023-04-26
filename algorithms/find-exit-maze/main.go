package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // inicializar el generador de números aleatorios

	// generar la cantidad aleatoria de filas y columnas
	filas := rand.Intn(16) + 5    // generará un número entre 5 y 20
	columnas := rand.Intn(16) + 5 // generará un número entre 5 y 20

	// crear la matriz de ceros
	matriz := make([][]int, filas)
	for i := range matriz {
		matriz[i] = make([]int, columnas)
	}

	// imprimir la matriz con el ancho y alto de filas y columnas
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Printf("%5d", matriz[i][j])
		}
		fmt.Println()
	}
}
