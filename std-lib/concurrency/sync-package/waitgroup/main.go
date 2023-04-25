package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from foo!")
	}
	wg.Done()
}

func bar(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from bar!")
	}
	wg.Done()
}

func main() {
	// TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()

	// also correct, but not relevant in this case
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	data++
	// }(&wg)

	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Hello from main!")

	wg.Add(2)
	go foo(&wg)
	go bar(&wg)

	wg.Add(3)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello from anonymous func 1!")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello from anonymous func 2!")
		}
		wg.Done()
	}()

	func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello from main goroutine!")
		}
	}()

	go foo(&wg)

	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Num of CPUs:", runtime.NumCPU())

	fmt.Println("Done..")

}
