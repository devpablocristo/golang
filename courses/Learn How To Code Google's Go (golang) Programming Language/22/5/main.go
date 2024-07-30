package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	fin := make(chan int)

	go enviar(par, impar, fin)

	recibir(par, impar, fin)

	fmt.Println("-----FIN-----")
}

func enviar(par, impar, fin chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	fin <- 0
}

func recibir(par, impar, fin <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el channel par:\t", v)
		case v := <-impar:
			fmt.Println("Desde el channel impar:\t", v)
		case v := <-fin:
			fmt.Println("Desde el channel fin:\t", v)
			return
		}
	}
}
