package main

import "fmt"

func main() {

	c := make(chan int)

	// enviar
	go foo(c)

	// recibir
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("-----FIN-----")

}

// enviar
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
