package jwtpkg

import (
	"fmt"
	"sync"

	"github.com/golang-jwt/jwt/v5"

	ports "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
)

var (
	jwtInstance ports.JwtClient
	jwtOnce     sync.Once
	initError   error
)

type jwtService struct {
	secretKey string
}

// newJWTService inicializa el servicio JWT con una clave secreta
func newJWTService(secretKey string) (ports.JwtClient, error) {
	jwtOnce.Do(func() {
		if secretKey == "" {
			initError = fmt.Errorf("secret key cannot be empty")
			return
		}

		jwtInstance = &jwtService{
			secretKey: secretKey,
		}
	})
	return jwtInstance, initError
}

// GetJWTInstance devuelve la instancia del cliente JWT
func GetJWTInstance() (ports.JwtClient, error) {
	if jwtInstance == nil {
		return nil, fmt.Errorf("JWT service is not initialized")
	}
	return jwtInstance, nil
}

// GenerateToken genera un token JWT con las reclamaciones proporcionadas
func (j *jwtService) GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken valida un token JWT proporcionado
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
