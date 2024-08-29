package main

import (
	"context"
	"log"

	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter-client"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkgrpcclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	"go-micro.dev/v4/logger"
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

	logger.Info(greeterUseCases.Greet(context.Background()))  
}

func setupServices() sdkgrpcclientport.Client {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return c
}
