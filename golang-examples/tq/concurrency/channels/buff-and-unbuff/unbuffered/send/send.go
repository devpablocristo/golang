package main

import "fmt"

func get(c chan string) {
	fmt.Println(c)
}

func main() {
	c := make(chan string)

	c <- "Hello"

	for i := 0; i < 10; i++ {
		go get(c)
	}
}
