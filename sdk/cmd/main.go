package main

import (
	"fmt"

	loggerpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	viperpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	gomicropkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	ginpkg "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
)

// NOTE: mover examples/go-micro
func main() {
	if err := viperpkg.LoadConfig("../"); err != nil {
		loggerpkg.StdError("Viper Service error: %v", err)
	}

	gomicroService, err := gomicropkg.Bootstrap()
	if err != nil {
		loggerpkg.StdError("GoMicro Service error: %v", err)
	}

	loggerpkg.GmInfo("Starting application...")

	fmt.Println("Registro de servicios:", gomicroService.GetRegistry())

	ginService, err := ginpkg.Bootstrap()
	if err != nil {
		loggerpkg.GmError("Gin Service error: %v", err)
	}

	r := ginService.GetRouter()

	_ = r

	go func() {
		if err := ginService.RunServer(); err != nil {
			loggerpkg.GmError("Failed to run Gin: %v", err)
		}
	}()

	if err := gomicroService.Start(); err != nil {
		loggerpkg.GmError("Error starting GoMicro Service: %v", err)
	}
}
