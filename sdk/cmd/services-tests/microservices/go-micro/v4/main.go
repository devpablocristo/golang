package main

import (
	"fmt"
	"log"

	gmbtrap "github.com/devpablocristo/golang/sdk/internal/bootstrap/go-micro"
	inisetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/initial"
	"github.com/spf13/viper"
)

func main() {
	if err := inisetup.SetupViperConfig("./config"); err != nil {
		log.Fatalf("Failed to set up Viper config: %v", err)
	}

	// Ahora puedes acceder a las configuraciones usando Viper
	fmt.Println("App Name:", viper.GetString("APP_NAME"))

	// Lanzar el servicio Go-Micro
	gmService, err := gmbtrap.BootstrapGoMicro()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Registro de servicios:", gmService.GetRegistry())

	// Iniciar el servicio
	if err := gmService.Start(); err != nil {
		log.Fatalf("Error al iniciar el servicio: %v", err)
	}
}
