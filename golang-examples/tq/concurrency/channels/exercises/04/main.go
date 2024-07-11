package main

func main() {

	// 1. Crear el portal
	portal := make(chan string)

	// 2. Abrir el portal al universo 2
	go universo2(portal)

	// 3. Abrir a portal en el universo 3
	universo3(portal)
	// 4. Enviar a los heroes!

}

func universo2(portal chan string) {
	portal <- "ironman"
	portal <- "thor"
	portal <- "spiderman"
}

func universo3(portal chan string) {

}
