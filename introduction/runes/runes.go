package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Olá, 世界"

	// 1. Longitud en bytes
	fmt.Println("Bytes:", len(s)) // Bytes: 11

	// 2. Número de runas (caracteres)
	fmt.Println("Runas:", utf8.RuneCountInString(s)) // Runas: 7

	// 3. Indexación directa (byte) vs conversión a []rune
	fmt.Printf("Primer byte: %x\n", s[0]) // Primer byte: c3
	runes := []rune(s)
	fmt.Printf("Primera runa: %c\n", runes[0]) // Primera runa: O

	// 4. Recorrido adecuado
	for i, r := range s {
		fmt.Printf("Runa %c en byte offset %d\n", r, i)
	}
}
