package main

import (
	"log"

	gtwgreeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter-server"
	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter-server"
	sdklogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server"
	sdkgrpcserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gServer := setupServices()

	sdklogger.Info("Estandar Info!")
	sdklogger.Mwarning("Este es un warning!")

	greeterUseCases := greeter.NewUseCases()
	greeterGrpcServer := gtwgreeter.NewGrpcServer(greeterUseCases, gServer)
	greeterGrpcServer.Start()

	if err := greeterGrpcServer.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func setupServices() sdkgrpcserverport.Server {
	s, err := sdkgrpcserver.Bootstrap()
	if err != nil {
		log.Fatalf("failed to initialize gRPC server: %v", err)
	}

	return s
}
