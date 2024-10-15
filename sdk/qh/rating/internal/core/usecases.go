package rating

import (
	"context"
	"fmt"

	entities "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/ports"
)

type useCases struct {
	grpcClient ports.GrpcClient
}

// NewUseCases crea una nueva instancia de useCases
func NewUseCases(gc ports.GrpcClient) ports.UseCases {
	return &useCases{
		grpcClient: gc,
	}
}

// Login maneja la lógica de autenticación de usuario
// func (s *useCases) Login(ctx context.Context, creds *entities.LoginCredentials) (*entities.Token, error) {
// 	userUUID, err := s.grpcClient.GetUserUUID(ctx, creds)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al obtener el UUID del usuario: %w", err)
// 	}

// 	token, err := s.jwtService.GenerateToken(userUUID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error generando el token de autenticación: %w", err)
// 	}

// 	// Devuelve el token generado
// 	return token, nil
// }
