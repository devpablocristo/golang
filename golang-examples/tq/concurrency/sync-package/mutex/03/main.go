package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		show(i) // synchronous call
	}

	duration := time.Since(start).Milliseconds()
	fmt.Println("Total time:", duration)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGoroutine(i, wg, m) // concurrent call
	}

	wg.Wait()
	duration = time.Since(start).Milliseconds()
	fmt.Println("Total time:", duration)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		show(i)                    // synchronous call
		go showGoroutine(i, wg, m) // concurrent call
	}

	wg.Wait()
	duration = time.Since(start).Milliseconds()

	fmt.Println("Total time:", duration)
}

func show(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Id %d: sleeping for %dms\n", id, delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func showGoroutine(id int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	delay := rand.Intn(500)
	fmt.Printf("Goroutine %d: sleeping for %dms\n", id, delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)

	m.Unlock()
	wg.Done()
}
