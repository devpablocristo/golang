package main

import (
	"log"

	gtwgreeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter"
	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter"
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

	greeterUseCases := greeter.NewUseCases(nil)
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
