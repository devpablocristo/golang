package main

import (
	"context"
	"log"
	"time"

	sdklogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkgrpcclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	calculator "github.com/devpablocristo/golang/sdk/services/calculator-client/internal/calculator-client"
)

func init() {
	//if err := sdkviper.LoadConfig("../../../"); err != nil { // en local
	if err := sdkviper.LoadConfig(); err != nil { // con docker

		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gClient := setupServices()

	calculatorGrpcClient := calculator.NewGrpcClient(gClient)

	calculatorUseCases := calculator.NewUseCases(calculatorGrpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// example
	firstNum := int32(23)
	lastNum := int32(12)

	res, err := calculatorUseCases.Addition(ctx, firstNum, lastNum)
	if err != nil {
		log.Fatalf("Error calling Greet method: %v", err)
	}
	sdklogger.Info("%d", res)
}

func setupServices() sdkgrpcclientport.Client {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	return c
}
