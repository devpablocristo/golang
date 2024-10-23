package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	sdkcnfldr "github.com/devpablocristo/golang/sdk/pkg/config/configLoader"

	companyconn "github.com/devpablocristo/golang/sdk/sg/users/internal/company/adapters/connectors"
	personconn "github.com/devpablocristo/golang/sdk/sg/users/internal/person/adapters/connectors"

	userconn "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors"
	usergtw "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/gateways"
	user "github.com/devpablocristo/golang/sdk/sg/users/internal/core"
)

func init() {
	if err := sdkcnfldr.LoadConfig("config/.env", "config/.env.local"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}

	fmt.Println("checking env 'AFIP_REALM':", viper.GetString("AFIP_REALM"))
}

func main() {

	usersRepo, err := userconn.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	personRepo, err := personconn.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	companyRepo, err := companyconn.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	usersUseCases := user.NewUseCases(usersRepo, personRepo, companyRepo)

	userHandler, err := usergtw.NewGinHandler(usersUseCases)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Iniciar el servidor de Gin
	err = userHandler.Start()
	if err != nil {
		log.Fatalf("Gin Server error at start: %v", err)
	}
}
