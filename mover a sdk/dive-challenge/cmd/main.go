package main

import (
	"log"

	api "github.com/devpablocristo/dive-challenge/cmd/api"
)

func main() {
	log.Println("Starting application...")

	config, err := api.Config()
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := api.Build(config)
	if err := app.Run(); err != nil {
		log.Fatal("error at app startup", err)
	}
}
