package greeter

import (
	"context"
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/services/greeter-client/internal/greeter-client/ports"
)

type useCases struct {
	grpcClient ports.GrpcClient
}

func NewUseCases(gc ports.GrpcClient) ports.UseCases {
	return &useCases{
		grpcClient: gc,
	}
}

func (c *useCases) Greet(ctx context.Context, firstName, lastName string) (string, error) {
	message, err := c.grpcClient.Greet(ctx, firstName, lastName)
	if err != nil {
		return "", fmt.Errorf("error calling gRPC method: %v", err)
	}

	return message, nil
}
