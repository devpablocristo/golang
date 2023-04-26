package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	ch := make(chan string)
	go generateID(ch)

	fmt.Println(<-ch)
}

func generateID(ch chan string) {
	id := uuid.New()
	ch <- id.String()
	close(ch)
}
