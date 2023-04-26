// go run -race 4.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inc := 0
	gs := 100

	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {

			m.Lock()

			v := inc
			v++
			inc = v
			fmt.Println(inc)

			m.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("inc:", inc)
}
