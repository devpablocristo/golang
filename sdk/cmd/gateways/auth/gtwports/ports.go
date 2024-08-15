package gtwports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type MessageBroker interface {
	GetUserUUID(context.Context, *entities.LoginCredentials) (string, error)
}

type GrpcClient interface {
	GetUserUUID(string, string) (string, error)
}
