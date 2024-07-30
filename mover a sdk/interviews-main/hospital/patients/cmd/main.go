package main

import (
	"log"
	"os"

	"github.com/devpablocristo/interviews/hospital2/patients/api"
)

const defaultPort = "8080"

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
	api.Start(port)
	//	other.Start(param)
}
