package main

import "fmt"

func main() {

	n := 15
	fmt.Println(reverseInt(n))

	n = 12345
	fmt.Println(reverseInt(n))

	n = -12345
	fmt.Println(reverseInt(n))
}

func reverseInt(n int) int {
	var res int
	var neg bool

	if n < 0 {
		n = -n
		neg = true
	}

	for n > 0 {
		res = res*10 + n%10
		n = n / 10
	}

	if neg {
		res = -res
	}

	return res

}
