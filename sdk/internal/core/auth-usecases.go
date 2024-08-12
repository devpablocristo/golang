package core

import (
	"context"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/ports"

	"github.com/devpablocristo/golang/sdk/internal/core/auth"
)

type AuthUseCases interface {
	Login(context.Context, string) (*auth.Auth, error)
}

type authUseCases struct {
	broker ports.MessageBroker
}

func NewAuthUseCases(b ports.MessageBroker) AuthUseCases {
	return &authUseCases{
		broker: b,
	}
}

func (s *authUseCases) Login(ctx context.Context, userUUID string) (*auth.Auth, error) {
	return nil, nil
}
