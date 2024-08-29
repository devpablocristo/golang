package main

import (
	"log"

	greeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	// Inicializar el cliente gRPC usando Bootstrap
	client, err := sdkgrpcclient.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}// Send requests to the server

	// Crear una instancia del adaptador GrpcClient
	greeterClient := greeter.NewGrpcClient(client)

	// Llamar a MyMethod para realizar la operación gRPC
	message, err := greeterClient.SayHello("World")
	if err != nil {
		log.Fatalf("Error calling gRPC method: %v", err)
	}

	// Imprimir el mensaje de la respuesta
	log.Printf("Response from gRPC server: %v", message)

	// Cerrar el cliente gRPC
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close gRPC client: %v", err)
	}
}
