package main

import (
	"log"

	gtwgreeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter"
	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkgrpcclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	sdkgrpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server"
	sdkgrpcserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gClient, gServer := setupServices()

	greeterGrpcClient := greeter.NewGrpcClient(gClient)
	greeterUseCases := greeter.NewUseCases(greeterGrpcClient)
	greeterGrpcServer := gtwgreeter.NewGrpcServer(greeterUseCases, gServer)
	greeterGrpcServer.Start()

	if err := greeterGrpcServer.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func setupServices() (sdkgrpcclientport.Client, sdkgrpcserverport.Server) {
	c, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	s, err := sdkgrpcserver.Bootstrap()
	if err != nil {
		log.Fatalf("failed to initialize gRPC server: %v", err)
	}

	return c, s
}
