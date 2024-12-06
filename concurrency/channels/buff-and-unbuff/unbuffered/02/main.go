package main

import "fmt"

func main() {
	portal := make(chan string)

	go universo2(portal)

	fmt.Println(<-portal)
	fmt.Println(<-portal)
	fmt.Println(<-portal)
}

func universo2(portal chan string) {
	portal <- "Ironman"
	portal <- "Thor"
	portal <- "Spiderman"
}
