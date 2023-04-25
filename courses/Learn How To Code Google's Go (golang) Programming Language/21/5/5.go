// go run -race 5.go

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var inc int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			fmt.Println(atomic.LoadInt64(&inc))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("inc:", inc)
}
