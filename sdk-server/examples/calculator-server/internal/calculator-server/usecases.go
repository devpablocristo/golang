package calculator

import (
	"context"

	"github.com/devpablocristo/golang/sdk/examples/calculator-server/internal/calculator-server/ports"
)

type useCases struct{}

func NewUseCases() ports.UseCases {
	return &useCases{}
}

func (arith useCases) Addition(ctx context.Context, a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith useCases) Subtraction(ctx context.Context, a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith useCases) Multiplication(ctx context.Context, a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith useCases) Division(ctx context.Context, a int32, b int32) (int32, error) {
	return a / b, nil
}
