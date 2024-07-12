package main

import "fmt"

func main() {
	// main go routine
	// universo 1
	// crear el portal
	portal := make(chan string)

	// universo2
	// abrir el portal al universo 2
	go universo2(portal)
	// abrir el portal en el unvierso 3
	universo3(portal)
}

// universo 2
// envia a los heroes!
func universo2(portal chan string) {
	portal <- "Ironman"
	portal <- "Thor"
	portal <- "Spiderman"
}

// universo 3
// recibe a los heroes!
func universo3(portal <-chan string) {
	fmt.Println(<-portal)
	fmt.Println(<-portal)
	fmt.Println(<-portal)
}
