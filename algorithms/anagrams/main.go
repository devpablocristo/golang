package main

import (
	"fmt"
	"sort"
)

func main() {
	a := "aabcd"
	b := "aefgha"

	// ordenar a y b
	rs1 := []rune(a)
	rs2 := []rune(b)

	// ordenar los slices de runes
	sort.Slice(rs1, func(i, j int) bool {
		return rs1[i] < rs1[j]
	})

	sort.Slice(rs2, func(i, j int) bool {
		return rs2[i] < rs2[j]
	})

	r1 := findDiff(rs1, rs2)
	r2 := findDiff(rs2, rs1)

	r3, c1 := removeDuplicate(r1)
	r4, c2 := removeDuplicate(r2)

	counter := len(rs1) + len(rs2) - c1 - c2

	fmt.Println(string(r3))
	fmt.Println(string(r4))
	fmt.Println(counter)

}

func findDiff(rs1, rs2 []rune) []rune {
	result := []rune{}
	founded := false
	var r rune

	for _, c1 := range rs1 {
		for _, c2 := range rs2 {
			if c1 == c2 {
				r = c1
				founded = true
				break
			}
		}
		if founded {
			result = append(result, r)
		}
		founded = false
	}

	return result
}

func removeDuplicate(s []rune) ([]rune, int) {
	seen := make(map[rune]struct{})
	result := []rune{}
	counter := 0

	for _, char := range s {
		if _, exists := seen[char]; !exists {
			seen[char] = struct{}{}
			result = append(result, char)
		} else {
			counter++
		}
	}

	return result, counter
}
