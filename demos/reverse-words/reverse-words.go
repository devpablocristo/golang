package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "today is the first day of the rest of my life"

	// Invertir palabras manualmente
	words := strings.Fields(s)
	reserved := words[0]

	for i := len(words) - 1; i > 0; i-- {
		reserved += " " + words[i]
	}

	fmt.Println("Manual Reverse Words:", reserved)

	// Invertir palabras usando reverseWords
	reversedWords := reverseWords(s)
	fmt.Println("Reverse Words (Function):", reversedWords)

	// Invertir caracteres usando reverseZeroIdAlgo
	reversedChars := reverseZeroIdAlgo(s)
	fmt.Println("Reverse Characters (Rune Method):", reversedChars)

	// Invertir caracteres usando reverse
	reversed := reverse(s)
	fmt.Println("Reverse Characters (Accumulative):", reversed)
}

// Invertir el orden de las palabras en una frase
func reverseWords(input string) string {
	values := strings.Fields(input)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return strings.Join(values, " ")
}

// Invertir los caracteres de una cadena usando runas
func reverseZeroIdAlgo(s string) string {
	r := []rune(s)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

// Invertir los caracteres de una cadena acumulándolos
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
