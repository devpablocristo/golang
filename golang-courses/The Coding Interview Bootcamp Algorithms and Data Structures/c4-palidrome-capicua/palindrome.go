package main

import "fmt"

func main() {
	pal := "abba"
	noPal := "nirvana"

	// pal := "queen"
	// noPal := "racecar"

	check := false
	if pal == reverseString(pal) {
		check = true
	}
	fmt.Println("Is ", pal, " a palindrome: ", check)

	check = false
	if noPal == reverseString(noPal) {
		check = true
	}
	fmt.Println("Is ", noPal, " a palindrome: ", check)

	fmt.Println("Is ", pal, " a palindrome: ", palindrome(pal))
	fmt.Println("Is ", noPal, " a palindrome: ", palindrome(noPal))

}

func reverseString(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}

func palindrome(s string) bool {
	p := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			p = false
			break
		}
	}
	return p
}
