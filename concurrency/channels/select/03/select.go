package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go send()
	wg.Wait()

	fmt.Println("About to exit.")

}

func send() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)

	}
	wg.Done()
}
