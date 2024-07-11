package main

import "fmt"

// Implement relaying of message with Channel Direction

// Example:
// func pong (in <-chan string, out chan<- string) {}
// in <-chan, is a recibe only channel
// out chan<-, is a send only channel

func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "hi"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	v := <-ch1
	// send it on ch2
	ch2 <- v + " and bye"
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)

	// recv message on ch2
	go relayMsg(ch1, ch2)
	fmt.Println(<-ch2)
}
