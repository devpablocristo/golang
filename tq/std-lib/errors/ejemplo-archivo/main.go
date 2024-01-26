package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func mayusculas(s string) (string, error) {
	if s == "" {
		return "", errors.New("string vacio!")
	}
	return strings.ToTitle(s), nil
}

func archivo() {

	/*
		Formatos de impresión de errores:
			fmt.Println()
			log.Println()
			log.Fatalln()
			os.Exit()
			log.Panicln()
			panic()
	*/

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("log.Println - Ocurrió un error: ", err)
		return
	}

}

func main() {
	s, err := mayusculas("hola")
	if err != nil {
		fmt.Println("No se pudo cambiar a mayusculas: ", err)
		return
	}

	fmt.Println("Exito!: ", s)

	s, err = mayusculas("")
	if err != nil {
		fmt.Println("No se pudo cambiar a mayusculas: ", err)
		return
	}

	fmt.Println("Exito!: ", s)

}
