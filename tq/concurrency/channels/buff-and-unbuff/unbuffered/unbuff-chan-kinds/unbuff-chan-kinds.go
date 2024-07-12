package main

import "fmt"

// Example:
// func pong (in <-chan string, out chan<- string) {}
// in <-chan, is a recibe only channel
// out chan<-, is a send only channel
func saySomething(c chan string) {
	v := <-c
	c <- v + " World!"
}

func sendMsg(c chan<- string) {
	c <- "Fortunata!"
	close(c)
}

func recibeMsg(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	// bidireccinal unbuffered channel
	bidDirChan := make(chan string)
	// to avoid deadlocks, ALWAYS the channel must be listing, before hearing
	go saySomething(bidDirChan)
	bidDirChan <- "Hello"
	msg, _ := <-bidDirChan
	fmt.Println(msg)
	close(bidDirChan)

	// send only unbuffered channel
	sndOnlyChan := make(chan string)
	go sendMsg(sndOnlyChan)
	fmt.Println(<-sndOnlyChan)

	// recibe only unbuffered channel
	recOnlyChan := make(chan string)
	// to avoid deadlocks, ALWAYS the channel must be listing, before hearing
	go recibeMsg(recOnlyChan)
	recOnlyChan <- "Toribio!"
	close(recOnlyChan)

}
