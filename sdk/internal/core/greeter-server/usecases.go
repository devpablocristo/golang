package greeter

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter-server/ports"
)

type useCases struct {
}

func NewUseCases() ports.UseCases {
	return &useCases{}
}

func (c *useCases) SayHello(ctx context.Context) (string, error) {
	message := "Hi!"

	return message, nil
}
