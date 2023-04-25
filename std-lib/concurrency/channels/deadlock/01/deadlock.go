package main

import "fmt"

func sum(n1 int, n2 int, chsum chan int) {
	sum := n1 + n2
	chsum <- sum
}

func main() {
	ch := make(chan int)
	//go sum(2, 5, ch)  // correct
	sum(2, 5, ch) // incorrect
	result := <-ch
	fmt.Println(result)
}
