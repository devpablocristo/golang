// How to stop a goroutine

package main

import "fmt"

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("The goroutine will stop now")
				return
			default:
				// …
			}
		}
	}()
	// …
	quit <- true
}
