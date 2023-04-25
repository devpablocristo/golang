package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var incrementer int64 = 0

	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()

		fmt.Println("Nº gouroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println(incrementer)
}
