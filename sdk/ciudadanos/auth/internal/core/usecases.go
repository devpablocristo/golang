package authe

import (
	"context"
	"fmt"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"

	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
)

type useCases struct {
	jwtService ports.JwtService
}

func NewUseCases(js ports.JwtService) ports.UseCases {
	return &useCases{
		jwtService: js,
	}
}

// Login maneja la lógica de autenticación de usuario
func (s *useCases) Login(ctx context.Context, creds *sdktypes.LoginCredentials) (*sdkjwt.Token, error) {
	// userUUID, err := s.grpcClient.GetUserUUID(ctx, creds)
	// if err != nil {
	// 	return nil, fmt.Errorf("error al obtener el UUID del usuario: %w", err)
	// }

	token, err := s.jwtService.GenerateToken("userUUID")
	if err != nil {
		return nil, fmt.Errorf("error generando el token de autenticación: %w", err)
	}

	// Devuelve el token generado
	return token, nil

}
