package main

import (
	"os"

	"github.com/devpablocristo/interviews/bookstore/apps/inventory/backend/api"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
	//	other.Start(param)
}
