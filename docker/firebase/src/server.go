package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//investigar como se usa url.path
	fmt.Fprintf(w, "Hello World!: %s", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	fmt.Println("Server listening!")

	http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
