package greeter

import (
	"context"
	"log"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter/ports"
)

type useCases struct {
	grpcClient ports.GrpcClient
}

func NewUseCases(gc ports.GrpcClient) ports.UseCases {
	return &useCases{
		grpcClient: gc,
	}
}

func (c *useCases) Hello(ctx context.Context) (string, error) {
	message, err := ports.GrpcClient.SayHello("World")
	if err != nil {
		log.Fatalf("Error calling gRPC method: %v", err)
	}

	return message, nil
}
