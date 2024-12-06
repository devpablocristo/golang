package main

import "fmt"

func main() {
	portal := make(chan string)

	go fmt.Println("hola")

	portal <- "Ironman"

	fmt.Println(<-portal)
}
