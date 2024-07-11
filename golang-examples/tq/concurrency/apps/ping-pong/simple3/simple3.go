package main

import (
	"fmt"
	"sync"
)

func readUni(m <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-m)
	wg.Done()
}

func main() {
	rmsg := make(chan string, 1)
	wg := sync.WaitGroup{}
	rmsg <- "ping pong"
	wg.Add(1)
	go readUni(rmsg, &wg)
	wg.Wait()
	close(rmsg)
}
