package authports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/entities"
)

// UseCases define las operaciones de casos de uso para autenticación
type UseCases interface {
	Login(ctx context.Context, creds *entities.LoginCredentials) (*entities.Token, error)
}

// JwtService define las operaciones del servicio JWT
type JwtService interface {
	GenerateToken(userUUID string) (*entities.Token, error)
}

// GrpcClient define las operaciones del cliente gRPC
type GrpcClient interface {
	GetUserUUID(ctx context.Context, cred *entities.LoginCredentials) (string, error)
}
