package coreauthports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type AuthUseCases interface {
	Login(context.Context, *entities.LoginCredentials) (*entities.Token, error)
}

type AccessControl interface {
	GenerateToken(map[string]any) error
}
