package portscore

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type AuthUseCases interface {
	Login(context.Context, *entities.LoginCredentials) (*entities.Token, error)
}

type AccessControl interface {
	GenerateToken(map[string]interface{}) error
}
