package main

import (
	"log"

	gtwcalculator "github.com/devpablocristo/golang/sdk/examples/calculator-server/gateways/calculator-server"
	calculator "github.com/devpablocristo/golang/sdk/examples/calculator-server/internal/calculator-server"
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

	calculatorUseCases := calculator.NewUseCases()
	calculatorGrpcServer := gtwcalculator.NewGrpcServer(calculatorUseCases, gServer)
	calculatorGrpcServer.Start()

	if err := calculatorGrpcServer.Start(); err != nil {
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
