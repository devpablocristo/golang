package main

import "fmt"

func main() {
	mazo := nuevoMazo()

	mazo.mezclar()
	mazo.mostrar()

	mano, restoDelMazo := repartir(3, mazo)
	mano.mostrar()
	fmt.Println("**********************")
	restoDelMazo.mostrar()

	//_ = mazo.guardarEnArchivoTxt("cartas.txt")
	//_ = mazo.leerArchivoTxt("cartas.txt")

	_ = mazo.guardarEnArchivoCsv("cartas")
	_ = mazo.leerArchivoCsv("cartas")

	nMazo := nuevoMazoDesdeArchivoCsv("cartas")
	nMazo.mostrar()
}
