package main

import "fmt"

// function receive two arg and return one
func plus(a int, b int) int {
	return a + b
}

// function receive two arg and return two
func hello(h string, w string) (string, string) {
	return h, w
}

func main() {
	fmt.Println("plus", plus(2, 6))
	fmt.Println(hello("hello", "golang"))
}
