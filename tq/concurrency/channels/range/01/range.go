package main

import "fmt"

func main() {
	ch := make(chan int, 6)

	go func() {
		// when the goroutine has sent all the values,
		// it needs to close the channel or there will be a deadlock
		defer close(ch)
		for i := 0; i < 6; i++ {
			// send iterator over channel
			ch <- i
		}
	}()

	// range over channel to recieve values
	for v := range ch {
		fmt.Println(v)
	}

}
