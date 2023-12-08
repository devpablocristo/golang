package main

import (
	"log"
	"net/http"

	handler "compile-daemon/infra/handler"
)

func main() {
	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/users", handler.UserPage)
	// puerto del contenedor
	log.Fatal(http.ListenAndServe(":8888", nil))
}
