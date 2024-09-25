package auth

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/entities"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type useCases struct {
	grpcClient   ports.GrpcClient
	redisService ports.RedisService
}

// NewUseCases crea una nueva instancia de useCases
func NewUseCases(gc ports.GrpcClient, rd ports.RedisService) ports.UseCases {
	return &useCases{
		grpcClient:   gc,
		redisService: rd,
	}
}

// Login maneja la lógica de autenticación de usuario
func (s *useCases) Login(ctx context.Context, creds *entities.LoginCredentials) (string, error) {
	// Devuelve el token generado
	return "token", nil
}
