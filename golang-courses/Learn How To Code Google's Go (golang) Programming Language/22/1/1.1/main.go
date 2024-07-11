package main

import "fmt"

func main() {
	// c es un channel
	// c := make(chan int, 1) // con este funca

	c := make(chan int)

	// pone 42 en c
	c <- 42
	// c se bloquea

	// saca lo que hay en c
	fmt.Println(<-c)

	/*
		este programa no fucinará
		fatal error: all goroutines are asleep - deadlock!
	*/
}
