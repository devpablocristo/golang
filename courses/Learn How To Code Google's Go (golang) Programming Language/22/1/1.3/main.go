package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)

	// no funcionará
	// se bloquea  también
	// fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-c) // con este no se bloquea
}
