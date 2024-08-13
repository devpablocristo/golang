package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/pkgports"
)

type useAuthCases struct {
	broker    gtwports.MessageBroker
	jwtClient pkgports.JWTClient
}

func NewAuthUseCases(b gtwports.MessageBroker, j pkgports.JWTClient) *useAuthCases {
	return &useAuthCases{
		broker:    b,
		jwtClient: j,
	}
}

func (s *useAuthCases) Login(ctx context.Context, user *entities.AuthUser) (*entities.Token, error) {
	_, err := s.broker.GetUserUUID(ctx, user)
	if err != nil {
		return nil, err
	}

	// Crear las declaraciones del token JWT
	claims := map[string]interface{}{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiración en 24 horas
	}

	// Generar el token usando el cliente JWT
	signedToken, err := s.jwtClient.GenerateToken(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Devolver el token generado
	return &entities.Token{
		AccessToken: signedToken,
		ExpiresAt:   time.Now().Add(time.Hour * 24),
	}, nil
}
