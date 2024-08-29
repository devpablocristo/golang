package greeter

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter/ports"
)

type useCases struct{}

func NewUseCases() ports.UseCases {
	return &useCases{}
}

func (s *useCases) Hello(ctx context.Context) (string, error) {
	return "hello", nil
}
