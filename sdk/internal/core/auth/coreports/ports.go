package coreports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type AuthUseCases interface {
	Login(context.Context, *entities.AuthUser) (*entities.Token, error)
}
