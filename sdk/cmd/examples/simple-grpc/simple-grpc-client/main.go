package main

import (
	"log"

	gtwgreeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter"
	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkgrpcclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gClient := setupServices()

	greeterGrpcClient := greeter.NewGrpcClient(gClient)
	greeterUseCases := greeter.NewUseCases(greeterGrpcClient)
	greeterGrpcServer := gtwgreeter.NewGrpcServer(greeterUseCases, nil)
	greeterGrpcServer.Start()

	if err := greeterGrpcServer.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func setupServices() sdkgrpcclientport.Client {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return c
}
