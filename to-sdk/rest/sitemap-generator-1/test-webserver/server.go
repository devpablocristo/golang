package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./test-webserver/html"))
	http.Handle("/", fs)

	port := ":8081"
	fmt.Println("Server is running on port" + port)
	http.ListenAndServe(":8081", nil)

}
