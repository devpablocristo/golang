package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/portscore"
	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/portspkg"
)

type useAuthCases struct {
	messageBroker gtwports.MessageBroker
	jwtClient     portspkg.JWTClient
}

func NewAuthUseCases(mb gtwports.MessageBroker, jc portspkg.JWTClient) portscore.AuthUseCases {
	return &useAuthCases{
		messageBroker: mb,
		jwtClient:     jc,
	}
}

func (s *useAuthCases) Login(ctx context.Context, user *entities.AuthUser) (*entities.Token, error) {
	_, err := s.messageBroker.GetUserUUID(ctx, user)
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
