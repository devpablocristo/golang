package main

import (
	"fmt"
	"sync"
)

func addOneHundred(message chan<- int, value int) {
	message <- value + 100
}

func main() {
	channel := make(chan int, 3) // create channel
	var wg sync.WaitGroup

	go addOneHundred(channel, 100) // send value 100 to channel
	go addOneHundred(channel, 200) // send value 200 to channel
	go addOneHundred(channel, 300) // send value 200 to channel

	v1 := <-channel // read channel
	v2 := <-channel // read channel
	v3 := <-channel // read channel

	fmt.Println("Value", v1)
	fmt.Println("Value", v2)
	fmt.Println("Value", v3)

	channel <- 1
	channel <- 1000
	channel <- 929292

	//wg.Add(1)
	go func() {
		for v := range channel {
			fmt.Println(v)
		}
		close(channel)
		wg.Done()
	}()
	//wg.Wait()

	//fmt.Scanln()
}
