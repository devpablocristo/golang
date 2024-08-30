package main

import (
	"context"
	"fmt"
	"log"

	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter-client"
	sdklogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
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

	fmt.Println("Holassss")

	greeterGrpcClient := greeter.NewGrpcClient(gClient)
	greeterUseCases := greeter.NewUseCases(greeterGrpcClient)

	sdklogger.Info(greeterUseCases.Greet(context.Background()))
}

func setupServices() sdkgrpcclientport.Client {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return c
}
