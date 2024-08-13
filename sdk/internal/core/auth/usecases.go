package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type useAuthCases struct {
	broker    gtwports.MessageBroker
	secretKey string // Llave secreta para firmar el token JWT
}

func NewAuthUseCases(b gtwports.MessageBroker, k string) *useAuthCases {
	return &useAuthCases{
		broker:    b,
		secretKey: k,
	}
}

func (s *useAuthCases) Login(ctx context.Context, user *entities.AuthUser) (*entities.Token, error) {
	_, err := s.broker.GetUserUUID(ctx, user)
	if err != nil {
		return nil, err
	}

	// Crear las declaraciones del token JWT
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiración en 24 horas
	}

	// Crear el token con las declaraciones
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la llave secreta
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	// Devolver el token generado
	return &entities.Token{
		AccessToken: signedToken,
		ExpiresAt:   time.Now().Add(time.Hour * 24),
	}, nil
}
