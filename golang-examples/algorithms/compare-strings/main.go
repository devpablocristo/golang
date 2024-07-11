package main

import (
	"fmt"
	"strings"
)

func main() {
	exact()
	partial()
	prefSuf()
	caseInsitive()
}

func exact() {
	palabras := []string{"hola", "mundo", "golang"}
	palabraBuscar := "golang"

	for _, palabra := range palabras {
		if palabra == palabraBuscar {
			fmt.Println("Encontrada:", palabra)
			break
		}
	}
}

func partial() {
	palabras := []string{"hola", "mundo", "golang"}
	substring := "go"

	for _, palabra := range palabras {
		if strings.Contains(palabra, substring) {
			fmt.Println("Encontrada:", palabra)
		}
	}
}

func prefSuf() {
	palabras := []string{"hola", "mundo", "golang"}
	prefijo := "go"

	for _, palabra := range palabras {
		if strings.HasPrefix(palabra, prefijo) {
			fmt.Println("Prefijo Encontrado:", palabra)
		}
	}
}

func caseInsitive() {
	palabras := []string{"Hola", "Mundo", "Golang"}
	palabraBuscar := "golang"

	for _, palabra := range palabras {
		if strings.EqualFold(palabra, palabraBuscar) {
			fmt.Println("Encontrada (sin importar mayúsculas/minúsculas):", palabra)
		}
	}
}
