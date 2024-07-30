package main

import "fmt"

func main() {
	c := make(<-chan int, 2)

	fmt.Printf("%T", c)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
