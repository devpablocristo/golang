package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	gs := 100

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(gs)

	incrementer := 0
	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println(runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(incrementer)

}
