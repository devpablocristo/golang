package greeter

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/examples/greeter-server/internal/core/greeter-server/ports"
)

type useCases struct {
}

func NewUseCases() ports.UseCases {
	return &useCases{}
}

func (c *useCases) Greet(ctx context.Context, name string) (string, error) {
	message := "Hi " + name + "!"

	return message, nil
}
