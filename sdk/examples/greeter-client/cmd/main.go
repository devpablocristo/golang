package main

import (
	"context"
	"log"
	"time"

	sdklogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkgrpcclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	greeter "github.com/devpablocristo/golang/sdk/services/greeter-client/internal/greeter-client"
)

func init() {
	//if err := sdkviper.LoadConfig("../../../"); err != nil { // en local
	if err := sdkviper.LoadConfig(); err != nil { // con docker

		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gClient := setupServices()

	log.Println("Holassssxxx")

	greeterGrpcClient := greeter.NewGrpcClient(gClient)

	greeterUseCases := greeter.NewUseCases(greeterGrpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// example
	firstName := "Flibaaais<2"
	lastName := "Bufé"

	res, err := greeterUseCases.Greet(ctx, firstName, lastName)
	if err != nil {
		log.Fatalf("Error calling Greet method: %v", err)
	}
	sdklogger.Info(res)
}

func setupServices() sdkgrpcclientport.Client {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	return c
}
