package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "2342452234456789976662323"
	s = Revrot(s, 4)
	fmt.Println(s)
}

func Revrot(s string, n int) string {

	if n > 0 {
		s = rotarInvertir(s, n)
	} else {
		return ""
	}
	return s
}

func sumaDeCubos(s string) int {
	sum := 0
	n := 0
	for i := 0; i < len(s); i++ {
		n, _ = strconv.Atoi(string(s[i]))
		sum += int(n) * int(n) * int(n)
	}
	return sum
}

func rotarInvertir(s string, n int) string {

	st := ""
	m := make(map[int]string)
	i := 0
	sum := 0

	for len(s) >= n {
		m[i] = s[:n]
		sum = sumaDeCubos(m[i])
		if sum%2 == 0 {
			fmt.Println("invertir:", m[i])
			m[i] = invertir(m[i])

		} else {
			fmt.Println("Rotar", m[i])
			m[i] = rotar(m[i])
		}
		s = s[n:]
		st += m[i]
		i++
	}
	return st
}

func invertir(s string) string {
	runas := []rune(s)
	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}
	return string(runas)
}

func rotar(s string) string {
	s1 := s[1:] + string(s[0])
	return s1
}
