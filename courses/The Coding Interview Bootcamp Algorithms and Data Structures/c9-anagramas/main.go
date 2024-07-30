package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

func main() {
	//Colinas – Nicolás	Calor – Carla
	//s1 := "rail safrety" // no es
	//s1 := "rail safety" // si es
	//s2 := "fairy tales"

	s1 := "RAIL!! SAFet@Y"
	s2 := "fairy 12312 ^  ..,,,  taleS"

	cleanStr1 := soloLetras(s1)
	cleanStr2 := soloLetras(s2)

	fmt.Println(soloLetras(cleanStr1))
	fmt.Println(soloLetras(cleanStr2))

	if anagramas(cleanStr1, cleanStr2) {
		fmt.Println("Es un anagrama.")
	} else {
		fmt.Println("NO es un anagrama.")

	}
}

// anagrama: mismas letras y misma cantidad de letras
func anagramas(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	anagrama := false

	bs1 := []byte(s1)
	bs2 := []byte(s2)

	if len(s1) == len(s2) {
		//fmt.Println("Tienen la misma cantidad de letra.")
		for _, v := range bs1 {
			m1[v] += 1
		}
		for _, v := range bs2 {
			m2[v] += 1
		}
		if reflect.DeepEqual(m1, m2) {
			anagrama = true
		}
	}
	return anagrama
}

func soloLetras(check string) string {
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal("Erro!", err)
	}
	lower := strings.ToLower(check)
	newStr := reg.ReplaceAllString(lower, "")
	return newStr
}
