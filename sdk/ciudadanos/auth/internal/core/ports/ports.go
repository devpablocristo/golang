package authports

import (
	"context"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"

	entities "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/entities"
)

// UseCases define las operaciones de casos de uso para autenticación
type UseCases interface {
	Login(context.Context, *sdktypes.LoginCredentials) (*sdkjwt.Token, error)
}

// JwtService define las operaciones del servicio JWT
type JwtService interface {
	GenerateToken(string) (*sdkjwt.Token, error)
}

type Repository interface {
	CreateEvent(context.Context) (*entities.Auth, error)
}
