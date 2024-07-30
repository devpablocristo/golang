package main

import "fmt"

func main() {
	//c := make(chan int, 2)
	c := make(<-chan int, 2) // solo recibe

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)

	/*
		no funciona, solo para recibir no para enviar
		invalid operation: <-c (receive from send-only type chan<- int)
	*/
}
