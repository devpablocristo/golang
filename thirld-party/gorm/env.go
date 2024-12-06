package main

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func loadEnv(dir string) error {
	err := godotenv.Load(dir + "/.env")
	if err != nil {
		log.Printf("unable to load .env file")
		return err
	}

	return nil
}
