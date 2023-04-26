package main

import (
	"log"
	"net/http"

	handler "github.com/devpablocristo/golang-docker-tutorial/handler"
)

func main() {
	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/users", handler.UserPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
