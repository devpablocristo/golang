package main

import "fmt"

func main() {
	c := make(chan int)

	cs := make(chan<- int) //send channel
	cr := make(<-chan int) // recibe channel

	go func() {
		c <- 43
		cs <- 21
	}()
	fmt.Println(<-c)

	//fmt.Println(<-cs)

	fmt.Print("------\n")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

}
