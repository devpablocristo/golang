package main

import (
	"fmt"
	"time"
)

func calculateFast(ch chan<- int, value int) {
	time.Sleep(3 * time.Second) // simulate 3 second of processing
	ch <- value * 10
}

func calculateSlow(ch chan<- int, value int) {
	time.Sleep(5 * time.Second) // simulate 5 second of processing
	ch <- value * 10
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go calculateFast(c1, 10)
	go calculateSlow(c2, 20)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1) // 100
		case msg2 := <-c2:
			fmt.Println("received", msg2) // 200
		default:
			fmt.Println("nothing received yet")
		}
	}
}
