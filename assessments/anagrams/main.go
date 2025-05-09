package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	s1 := "aabcd"
	s2 := "aefgha"

	rs1 := []rune(s1)
	rs2 := []rune(s2)

	fmt.Println(rs1)
	fmt.Println(rs2)

	slices.SortFunc(rs1, func(a, b rune) bool {
		return a < b
	})

	fmt.Println(rs1)

}
