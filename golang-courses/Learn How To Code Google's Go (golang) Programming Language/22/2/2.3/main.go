package main

import "fmt"

func main() {
	c := make(chan int)    // bireccional
	cr := make(<-chan int) // recibe
	ce := make(chan<- int) // envía

	fmt.Println("------------------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("ce\t%T\n", ce)

	fmt.Println("------------------------")
	fmt.Println("Específico a general")
	fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Printf("c\t%T\n", (chan int)(ce))

	/*
		no funca
		cannot convert cr (type <-chan int) to type chan int
		cannot convert ce (type chan<- int) to type chan int
	*/
}
