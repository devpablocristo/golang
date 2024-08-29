package main

import (
	"log"

	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	server, err := sdkgrpcserver.Bootstrap()
	if err != nil {
		log.Fatalf("failed to initialize gRPC server: %v", err)
	}

	// Registrar el servicio gRPC implementado
	gsa := greeter.NewGrpcServer()
	server.RegisterService(&pb.Greeter_ServiceDesc, gsa)

	// Iniciar el servidor gRPC
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
