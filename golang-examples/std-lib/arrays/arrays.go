package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Position 0 ", a[0])
	fmt.Println("Values of a", a)

	// array with assignment
	fibonacci := [11]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	fmt.Println("Values of fibonacci", fibonacci)

	// array with 3 position and two assignment
	array := [3]int{0, 1}
	fmt.Println("Values of array", array)

	// array multi-dimensional
	var twoDimensionArray [2][3]int
	twoDimensionArray[0][0] = 1
	twoDimensionArray[0][1] = 2
	twoDimensionArray[0][2] = 3

	twoDimensionArray[1][0] = 4
	twoDimensionArray[1][1] = 5
	twoDimensionArray[1][2] = 6

	fmt.Println("Values of twoDimensionArray", twoDimensionArray)

}
