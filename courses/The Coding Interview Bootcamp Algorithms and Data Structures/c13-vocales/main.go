package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	s := "aa1pqw22eiii"
	voc := vocales(s)
	m := contarVocales(voc)
	fmt.Println(m)
}

func contarVocales(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range []byte(s) {
		m[string(v)] += 1
	}
	return m
}

func vocales(s string) string {
	reg, err := regexp.Compile("[^aeiuo]+") // solo vocales
	//reg, err := regexp.Compile("^[aeiuo]+") // solo no vocales
	if err != nil {
		log.Fatal("Erro!", err)
	}
	lower := strings.ToLower(s)
	newStr := reg.ReplaceAllString(lower, "")
	return newStr
}
