package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	i := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for j := 0; j < gs; j++ {
		go func() {
			v := i
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			i = v
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Contador:", i)
}
