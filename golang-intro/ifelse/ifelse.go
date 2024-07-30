package main

import (
	"fmt"
)

func main() {

	// common if statement
	s := "hi"
	if s == "hi" {
		fmt.Println("Hello")
	} else if s == "bye" {
		fmt.Println("See you")
	} else {
		fmt.Println("How’s it going")
	}

	// if statement with operation
	x := 10
	n := 180
	if v := x + n; v <= 100 {
		fmt.Println("value is less or equal than 100 ", v)
	} else if v <= 200 {
		fmt.Println("value is less or equal than 200 ", v)
	} else {
		fmt.Println("the value is greater than 200", v)
	}

	// if without else
	i := 8
	j := 4
	if i%j == 0 {
		fmt.Println("It is divisible")
	}
}
