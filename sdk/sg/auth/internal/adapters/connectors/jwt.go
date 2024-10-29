package authconn

import (
	"fmt"
	"time"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdkjwtdefs "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/defs"

	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

type JwtService struct {
	jwtService sdkjwtdefs.Service
}

func NewJwtService() (ports.JwtService, error) {
	js, err := sdkjwt.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("jwt bootstrap error: %w", err)
	}

	return &JwtService{
		jwtService: js,
	}, nil
}

func (j *JwtService) GenerateToken(userUUID string) (string, error) {
	token, err := j.jwtService.GenerateTokenForSubject(userUUID, 24*time.Hour)
	if err != nil {
		return "", fmt.Errorf("error al generar el token: %w", err)
	}

	return token, nil
}

func (j *JwtService) ValidateToken(token string) (*sdkjwtdefs.TokenClaims, error) {
	tokenClaims, err := j.jwtService.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	return tokenClaims, nil
}
