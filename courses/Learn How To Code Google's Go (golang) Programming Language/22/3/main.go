package main

import "fmt"

func main() {

	c := make(chan int)

	// enviar
	go foo(c)

	// recibir
	bar(c)

	fmt.Println("-----FIN-----")

}

// enviar
func foo(c chan<- int) {
	c <- 42
}

// recibir
func bar(c <-chan int) {
	fmt.Println(<-c)
}
