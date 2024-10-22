package auth

import (
	"context"
	"fmt"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"

	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

type useCases struct {
	jwtService     ports.JwtService
	repository     ports.Repository
	httpClient     ports.HttpClient
	sessionManager ports.SessionManager
}

func NewUseCases(js ports.JwtService, rp ports.Repository, hc ports.HttpClient, sm ports.SessionManager) ports.UseCases {
	return &useCases{
		jwtService:     js,
		repository:     rp,
		httpClient:     hc,
		sessionManager: sm,
	}
}

// Login maneja la lógica de autenticación de usuario
func (u *useCases) Login(ctx context.Context, creds *sdktypes.LoginCredentials) (*sdkjwt.Token, error) {
	// userUUID, err := s.grpcClient.GetUserUUID(ctx, creds)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get user UUID: %w", err)
	// }

	token, err := u.jwtService.GenerateToken("userUUID")
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Return the generated token
	return token, nil
}

func (u *useCases) AfipLogin(ctx context.Context, jwtToken string) error {
	// if err := u.sessionManager.JwtToSession(ctx, jwtToken, "afip-login"); err != nil {
	// 	return fmt.Errorf("failed to save JWT to session: %w", err)
	// }
	return nil
}
