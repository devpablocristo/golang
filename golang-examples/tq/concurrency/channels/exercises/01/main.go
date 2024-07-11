package main

import "fmt"

func main() {
	// 1.
	var c1 chan float64

	// 2.
	// Declaring and initilizing a RECEIVE-ONLY channel
	c2 := make(<-chan rune)

	// Declaring and initilizing a SEND-ONLY channel
	c3 := make(chan<- rune)

	// 3.
	c4 := make(chan int, 10)

	// 4.
	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
}
