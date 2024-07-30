package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func characters() {
	for i := 'a'; i < 't'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	go numbers()
	go characters()

	fmt.Scanln()
	fmt.Println("done..")
}
