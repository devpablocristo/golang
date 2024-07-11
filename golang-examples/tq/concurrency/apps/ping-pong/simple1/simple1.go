package main

import (
	"fmt"
	"sync"
)

func writeBi(m chan string) {
	m <- "ping pong"
}

func writeUni(m chan<- string) {
	m <- "ping pong"
}

func readUni(m <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-m)
	wg.Done()
}

func input(m chan<- string, wg *sync.WaitGroup) {
	m <- "tori"
	wg.Done()
}

func output(m <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-m)
	wg.Done()
}

func main() {
	wmsg := make(chan string)
	go writeUni(wmsg)
	fmt.Println(<-wmsg)

	rmsg := make(chan string, 1)
	wg := sync.WaitGroup{}
	rmsg <- "ping pong"
	wg.Add(1)
	go readUni(rmsg, &wg)
	wg.Wait()

	imsg := make(chan string)
	wg.Add(2)
	go input(imsg, &wg)
	go output(imsg, &wg)
	wg.Wait()

	tmsg := make(chan string)
	go writeBi(tmsg)
	fmt.Println(<-tmsg)
}
