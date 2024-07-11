package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//var e, o, q chan int

	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	wg.Add(1)
	go send(e, o, q)
	wg.Wait()

	recibe(e, o, q)

	fmt.Println("About to exit.")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
			fmt.Println("e")
		} else {
			o <- i
			fmt.Println("o")
		}

		fmt.Print(i)
	}
	close(e)
	close(o)

	wg.Done()

	q <- 0
}

func recibe(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From de even channel: ", v)
		case v := <-o:
			fmt.Println("From de odd channel: ", v)
		case v := <-e:
			fmt.Println("From de quit channel: ", v)
			return
		}
	}
}
