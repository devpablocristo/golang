package main

import (
	"fmt"
	"strconv"
)

func main() {

	n1 := 12345
	n2 := -98765
	n3 := 400

	fmt.Println(darVueltaInt(n1))
	fmt.Println(darVueltaInt(n2))
	fmt.Println(darVueltaInt(n3))

	fmt.Println(darVueltaInt2(n1))
	fmt.Println(darVueltaInt2(n2))
	fmt.Println(darVueltaInt2(n3))

}

func darVueltaInt(n int) int {
	signo := 1
	if n < 0 {
		signo = -1
		n *= -1
	}

	s := strconv.Itoa(n)
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	rev, _ := strconv.Atoi(string(bs))
	rev *= signo

	return rev
}

func darVueltaInt2(n int) int {
	new_int := 0

	signo := 1
	if n < 0 {
		signo = -1
		n *= -1
	}

	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}

	new_int *= signo
	return new_int
}
