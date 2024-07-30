package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	fin := make(chan bool)

	go enviar(par, impar, fin)

	recibir(par, impar, fin)

	fmt.Println("-----FIN-----")
}

// channel enviar
func enviar(par, impar chan<- int, fin chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(fin)
}

// channel recibir
func recibir(par, impar <-chan int, fin <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el channel par:\t", v)
		case v := <-impar:
			fmt.Println("Desde el channel impar:\t", v)
		case v, ok := <-fin:
			if !ok {
				fmt.Println("Desde coma ok\t", v, ok)
				return
			} else {
				fmt.Println("Desde coma ok\t", v)
			}
		}
	}
}
