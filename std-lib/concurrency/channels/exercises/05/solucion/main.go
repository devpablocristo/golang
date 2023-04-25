package main

import "fmt"

func main() {

	portal := make(chan string, 3)

	planeta2(portal)
	planeta3(portal)

}

func planeta2(portal chan string) {
	portal <- "Ironman"
	portal <- "Thor"
	portal <- "Spiderman"
}

func planeta3(portal chan string) {
	fmt.Println(<-portal)
	fmt.Println(<-portal)
	fmt.Println(<-portal)
}
