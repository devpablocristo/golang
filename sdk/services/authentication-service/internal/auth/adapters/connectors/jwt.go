package authconn

import (
	"fmt"
	"time"

	sdk "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
	entities "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/entities"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type jwtService struct {
	JwtService sdkports.Service
}

func NewJwtService() (ports.JwtService, error) {
	js, err := sdk.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("Bootstrap error: %w", err)
	}

	return &jwtService{
		JwtService: js,
	}, nil
}

func (j *jwtService) GenerateToken(userUUID string) (*entities.Token, error) {
	// Definir los claims del JWT
	claims := map[string]interface{}{
		"sub": userUUID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	// Llama al SDK para generar el token
	tokenString, err := j.JwtService.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	// Crea la entidad Token con el token generado
	token := &entities.Token{
		AccessToken: tokenString,
		ExpiresAt:   time.Now().Add(time.Hour * 24), // Ejemplo de expiración de 24 horas
	}

	return token, nil
}
