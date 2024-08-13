package portscore

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type AuthUseCases interface {
	Login(context.Context, *entities.LogingCredentials) (*entities.Token, error)
}
