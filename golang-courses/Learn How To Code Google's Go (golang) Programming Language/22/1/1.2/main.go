package main

import "fmt"

func main() {
	// c es un channel
	c := make(chan int)

	// pone 42 en c
	go func() {
		c <- 42
		// el canal se bloquea aqui tb
		// pero como tiene su propia goroutine
		// el flow puede cuentinuar
	}()

	// saca lo que hay en c
	fmt.Println(<-c)

	/*
			este programa si fucinará
		 	Pq con la funcion y go, se crea una nueva goroutine
	*/
}
