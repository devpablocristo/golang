package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	sdkcnfldr "github.com/devpablocristo/golang/sdk/pkg/config/configLoader"

	userconn "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors"
	usergtw "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/gateways"
	personconn "github.com/devpablocristo/golang/sdk/sg/users/internal/person/adapters/connectors"

	user "github.com/devpablocristo/golang/sdk/sg/users/internal/core"
	person "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core"
)

func init() {
	if err := sdkcnfldr.LoadConfig("config/.env", "config/.env.local"); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}

	fmt.Println("checking env 'AFIP_REALM':", viper.GetString("AFIP_REALM"))
}

// NOTE: no pude implementar wire todavia, dan errores que no entiendo, mirar mas adelante
func main() {

	personRepo, err := personconn.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	personUseCases := person.NewUseCases(personRepo)

	userRepo, err := userconn.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	usersUseCases := user.NewUseCases(userRepo, personUseCases)

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
