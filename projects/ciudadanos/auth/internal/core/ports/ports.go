package authports

import (
	"context"

	sdkjwt "github.com/devpablocristo/golang/sdk/jwt/v5"
)

// UseCases define las operaciones de casos de uso para autenticación
type UseCases interface {
	Login(context.Context, *sdkjwt.LoginCredentials) (*sdkjwt.Token, error)
}

// JwtService define las operaciones del servicio JWT
type JwtService interface {
	GenerateToken(string) (*sdkjwt.Token, error)
}
