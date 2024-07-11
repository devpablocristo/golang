package main

import (
	"fmt"
	"strings"
)

func cleanString(s string) string {
	var cs string
	for _, char := range s {
		if char == 'A' || char == 'B' {
			cs += string(char)
		}
	}
	return cs
}

func removeExtrasAB(s string) string {
	sb := []byte(s)

	for i := 0; i < len(sb)-1; {
		if sb[i] == sb[i+1] {
			sb = append(sb[:i], sb[i+1:]...)
		} else {
			i++
		}
	}

	return string(sb)
}

func countExtrasAB(s string) int {
	var j int
	for i := 0; i < len(s)-1; i++ {
		if (s[i] == 'A' && s[i+1] == 'A') || (s[i] == 'B' && s[i+1] == 'B') {
			j++
		}
	}
	return j
}

func cleanString2(s string) string {
	return strings.Map(func(r rune) rune {
		if r == 'A' || r == 'B' {
			return r
		}
		return -1
	}, s)
}

func removeExtrasAB2(s string) string {
	var sb strings.Builder

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			sb.WriteByte(s[i])
		}
	}
	sb.WriteByte(s[len(s)-1])

	return sb.String()
}

func main() {
	s := "abcdefghAijklmnoApqBrstuvwBxyzAawqwqeBBjjj<"
	cs := cleanString(s)
	fmt.Println("Cleaned String:", cs)

	d := countExtrasAB(cs)
	fmt.Println("Count of Extras AB:", d)

	rs := removeExtrasAB(cs)
	fmt.Println("String after removing extras AB:", rs)
}
