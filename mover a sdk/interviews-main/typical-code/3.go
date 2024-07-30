// Write a Go program to find the nth Fibonacci number.

package main

import "fmt"

//nth fibonacci number function
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(7))
}
