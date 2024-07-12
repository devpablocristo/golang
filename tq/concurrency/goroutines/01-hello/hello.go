package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// goroutine function call
	go fun("goroutine func")

	// goroutine with anonymous function
	go func() {
		fun("goroutine anonymous func")
	}()

	// goroutine with function value call
	funcVar := fun
	go funcVar("goroutine func as value")

	// wait for goroutines to end
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done..")

}
