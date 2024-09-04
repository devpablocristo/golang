package calculator

import (
	"context"
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/internal/core/calculator-client/ports"
)

type useCases struct {
	grpcClient ports.GrpcClient
}

func NewUseCases(gc ports.GrpcClient) ports.UseCases {
	return &useCases{
		grpcClient: gc,
	}
}

func (c *useCases) Addition(ctx context.Context, firstNum, secondNum int32) (int32, error) {
	message, err := c.grpcClient.Addition(ctx, firstNum, secondNum)
	if err != nil {
		return 0, fmt.Errorf("error calling gRPC method: %v", err)
	}

	return message, nil
}
