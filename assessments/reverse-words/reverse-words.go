package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"

	fmt.Println("Invertir palabras:", simpleReverseWords(s))
	fmt.Println("Invertir runas:", simpleReverseRunes(s))
	fmt.Println("Invertir acumulativo:", simpleReverseAcc(s))
}

// 1. Invertir palabras de forma sencilla
func simpleReverseWords(s string) string {
	words := strings.Fields(s) // Divide la frase en palabras :contentReference[oaicite:0]{index=0}
	var rev []string
	for i := len(words) - 1; i >= 0; i-- {
		rev = append(rev, words[i]) // Añade cada palabra al slice resultado
	}
	return strings.Join(rev, " ") // Une las palabras con espacios
}

// 2. Invertir caracteres (runas) para manejar Unicode
func simpleReverseRunes(s string) string {
	runes := []rune(s) // Convierte la cadena en slice de runas :contentReference[oaicite:1]{index=1}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Intercambia extremos
	}
	return string(runes)
}

// 3. Inversión acumulativa de caracteres
func simpleReverseAcc(s string) string {
	var result string
	for _, r := range s { // Recorre runa por runa
		result = string(r) + result // Pre-pend de cada runa
	}
	return result
}
