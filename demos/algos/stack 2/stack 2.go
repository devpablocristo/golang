package main

import (
	"fmt"
	"sync"
)

type stack struct {
	items []int
	mu    sync.Mutex
}

func main() {
	s := &stack{}
	wg := &sync.WaitGroup{}

	// Lanzar goroutines para push concurrente
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.push(i)
			fmt.Println("items en push:", s.items)
		}(i)
	}

	wg.Wait()

	fmt.Println("Final items before popping:", s.items)

	for i := 0; i < 5; i++ {
		item := s.pop()
		fmt.Println("pop:", item)
		fmt.Println("items en pop:", s.items)
	}

	fmt.Println("Final items:", s.items)

}

func (s *stack) push(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("push:", i)
	s.items = append(s.items, i)
}

func (s *stack) pop() int {
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return -1
	}

	// Extrae el último elemento
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}
