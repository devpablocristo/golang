package main

import "fmt"

func main() {
	portal := make(chan string, 3)

	portal <- "Ironman"
	portal <- "Thor"
	portal <- "Spiderman"

	fmt.Println(<-portal)
	fmt.Println(<-portal)
	fmt.Println(<-portal)

}
