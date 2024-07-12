package main

import (
	"fmt"
	"sync"
)

func main() {

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	incrementer := 0
	for i := 0; i < gs; i++ {

		go func() {
			v := incrementer
			fmt.Println(incrementer)
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)

}
