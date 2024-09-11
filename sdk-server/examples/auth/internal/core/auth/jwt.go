package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	entities "github.com/devpablocristo/golang/sdk/examples/auth/internal/core/auth/entities"
)

// TODO: usar jwt del sdk
// JwtService define la interfaz del servicio JWT
type JwtService interface {
	GenerateToken(userUUID string) (*entities.Token, error)
}

// jwtService es la estructura que implementa la interfaz JwtService
type jwtService struct{}

// NewJwtService crea una nueva instancia de jwtService
func NewJwtService() JwtService {
	return &jwtService{}
}

// GenerateToken genera un JWT para un usuario dado su UUID
func (j *jwtService) GenerateToken(userUUID string) (*entities.Token, error) {
	// Lógica para generar el token (e.g., JWT)
	tokenString, err := createJWT(userUUID)
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

// createJWT crea un JWT firmado
func createJWT(userUUID string) (string, error) {
	// Define los claims del JWT
	claims := jwt.MapClaims{
		"sub": userUUID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	// Crea el token con firma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
