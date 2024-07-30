package main

import (
	"log"
	"os"

	//"github.com/devpablocristo/interviews/hospital2/patients/api"
	chi "github.com/devpablocristo/patients-api/patients/infrastructure/driver/http/chi"
)

const defaultPort = "8080"

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	chi.Start(port)
	chi.Start(port)
	//	other.Start(param)
}
