package main

import "fmt"

func main() {
	// main go routine
	// universo 1
	portal := make(chan string, 3)

	// universo2
	go enviarHeroes(portal)
	recibirHeroes(portal)
}

// universo 2
// chan<- only send channel
func enviarHeroes(portal chan<- string) {
	portal <- "Ironman"
	portal <- "Thor"
	portal <- "Spiderman"
}

// universo 3
// <-chan only recibe channel
func recibirHeroes(portal <-chan string) {
	fmt.Println(<-portal)
	fmt.Println(<-portal)
	fmt.Println(<-portal)
}
