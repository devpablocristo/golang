package main

import "fmt"

func CleanString(s string) string {
	var cleaned string
	for _, r := range s {
		if r == 'A' || r == 'B' {
			cleaned += string(r)
		}
	}
	return cleaned
}

func ConsecutivePairs(s string) int {
	var pairs int
	for i := 0; i < len(s); i++ {
		if s[i] == s[i+1] {
			pairs++
			i++
		}
	}
	return pairs
}

func RemoveConsecutivePairs(s string) string {
	var cleaned string

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			i++
		} else {
			cleaned += string(s[i])
		}
	}

	return cleaned
}

func main() {
	s := "abcdefghAijklmnoApqBrstuvwBxyzAawqwqeBBjjj"

	cleaned := CleanString(s)
	pairs := ConsecutivePairs(cleaned)
	pairsRemoved := RemoveConsecutivePairs(cleaned)

	fmt.Println(cleaned)
	fmt.Println(pairs)
	fmt.Println(pairsRemoved)

}
