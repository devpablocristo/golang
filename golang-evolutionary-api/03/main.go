package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome to the home page!")
	case "POST":
		fmt.Fprintf(w, "Post to the home page!")
	case "PUT":
		fmt.Fprintf(w, "Put to the home page!")
	case "DELETE":
		fmt.Fprintf(w, "Delete the home page!")
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye guys!")
}
