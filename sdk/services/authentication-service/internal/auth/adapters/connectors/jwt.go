package authconn

import (
	"log"
	"time"

	sdk "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
	entities "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/entities"
)

// JwtService define la interfaz del servicio JWT
type JwtService interface {
	GenerateToken(userUUID string) (*entities.Token, error)
}

// jwtService es la estructura que implementa la interfaz JwtService
type jwtService struct {
	JwtService sdkports.Service // Usa el cliente JWT del SDK
}

// NewJwtService crea una nueva instancia de jwtService usando el cliente JWT del SDK
func NewJwtService() JwtService {
	js, err := sdk.Bootstrap() // Inicializa el cliente JWT desde el SDK
	if err != nil {
		log.Fatalf("JWT Service error: %v", err)
	}

	return &jwtService{
		JwtService: js,
	}
}

// GenerateToken genera un JWT para un usuario dado su UUID
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
