package main

import (
	"fmt"
)

func digitsCounter(n int) int {
	if n < 10 {
		fmt.Println(n)
		return 1
	}
	fmt.Println(n)
	return 1 + digitsCounter(n/10)
}

func main() {
	n := 126236

	digits := digitsCounter(n)

	fmt.Println("Number of digits: ", digits)
}
