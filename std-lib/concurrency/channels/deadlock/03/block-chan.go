package main

import "fmt"

func main() {
	c := make(chan int, 2)
	d := make(chan int)
	e := make(chan int, 1)

	c <- 43
	c <- 34
	//c <- 11

	e <- 3
	//e <- 4 si dentro de la misma goroutine no se puede exceder el tamaño del buffer
	go func() {
		d <- 123
		e <- 1
		e <- 2
	}()

	fmt.Println(<-d)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(<-e)
	fmt.Println(<-e)
}

/*
canales
buf

unbuf
se bloquean si no estaban abiertos en ambos lados (goroutine)


*/
