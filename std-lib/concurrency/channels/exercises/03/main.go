// Coding Exercise #4

// Create a goroutine named power() that has one parameter of type int
// calculates the square value of its parameter and then sends the result into a channel.
// In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.
// Print out the square values.

package main

import "fmt"

// only send channel
func power(n int, res chan<- int) {

	sqrt := n * n

	res <- sqrt
}

func main() {
	// only recibe channel
	res := make(chan int)

	for i := 1; i < 51; i++ {
		go power(i, res)
		fmt.Println(i, " - ", <-res)
	}

}
