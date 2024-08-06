package main

import (
	"log"

	monitoring "github.com/devpablocristo/golang/sdk/cmd/rest/monitoring/routes"
	user "github.com/devpablocristo/golang/sdk/cmd/rest/user/routes"

	gingonicsetup "github.com/devpablocristo/golang/sdk/internal/platform/gin"
	gmwsetup "github.com/devpablocristo/golang/sdk/internal/platform/go-micro-web"
	basesetup "github.com/devpablocristo/golang/sdk/pkg/base-setup"
)

func main() {
	if err := basesetup.BasicSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	basesetup.LogInfo("Application started with JWT secret key: %s", basesetup.GetJWTSecretKey())
	basesetup.MicroLogInfo("Starting application...")

	gomicro, err := gmwsetup.NewGoMicroInstance()
	if err != nil {
		basesetup.MicroLogError("error initializing Go Micro: %v", err)
	}

	gingonic, err := gingonicsetup.NewGinInstance()
	if err != nil {
		basesetup.MicroLogError("error initializing Gin: %v", err)
	}

	monitoring.Routes(gingonic, gomicro)

	r := gingonic.GetRouter()

	user.Routes(r)

	gomicro.GetService().Handle("/", r)

	// Ejecuta Gin en la dirección especificada por Go-Micro
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run Gin: %v", err)
		}
	}()

	if err := gomicro.GetService().Run(); err != nil {
		basesetup.MicroLogError("error starting GoMicro service: %v", err)
	}
}
