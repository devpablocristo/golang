package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	sdkcnfldr "github.com/devpablocristo/golang/sdk/pkg/config/configLoader"
)

func init() {
	if err := sdkcnfldr.LoadConfig("config/.env", "config/.env.local"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}

	fmt.Println("checking env 'AFIP_REALM':", viper.GetString("AFIP_REALM"))
}

func main() {

	authHandler, err := InitializeApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Iniciar el servidor de Gin
	err = authHandler.Start()
	if err != nil {
		log.Fatalf("Gin Server error at start: %v", err)
	}
}