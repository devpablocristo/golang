package main

import "fmt"

func calculate(ch chan int, value int) {
	ch <- value * 10 // multiply value 10 into channel
}

func main() {

	// Get the value computed from goroutine
	ch := make(chan int)

	// send ch <- v
	// recibe v = <-ch
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	c, ok := <-ch
	// ok = true, value generated by write
	// ok = false, value generated by close
	if ok {
		fmt.Println("VOpen channel", c)
	}

	close(ch)

	c, ok = <-ch
	if !ok {
		fmt.Println("Closed channel", c)
	}

	ch2 := make(chan int) // create channel

	go calculate(ch2, 5)   // 50
	go calculate(ch2, 10)  // 100
	go calculate(ch2, 100) //1000

	v1 := <-ch2 // read ch2
	v2 := <-ch2 // read ch2

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)
	fmt.Println("<-channel:", <-ch2)

	close(ch2)

	fmt.Println("\ndone")
}
