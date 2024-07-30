package main

import "fmt"

func main() {

	q := make(chan int)
	c := gen(q)

	recibe(q, c)

	fmt.Println("About to exit.")

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
			fmt.Print(i)
		}
		close(c)
	}()

	return c
}

func recibe(q, c <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
