package main

import "fmt"

func foo(ch chan string) {
	defer close(ch)

	ch <- "hello"
	ch <- "bye"
	ch <- "toribio"
	ch <- "fufu"
}

func main() {
	c := make(chan int, 6)
	ch := make(chan string)

	//fmt.Println(<-ch)

	go foo(ch)

	for s := range ch {
		fmt.Println(s)
	}

	// you can read from closed channels
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// you can't write on closed channls
	//ch <- "coco"

	c <- 42
	c <- 43
	c <- 22
	c <- 5
	c <- 664
	c <- 6776

	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
