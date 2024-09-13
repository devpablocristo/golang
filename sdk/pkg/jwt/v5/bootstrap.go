package sdkjwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
)

// Bootstrap inicializa el servicio JWT usando una clave secreta generada
func Bootstrap() (ports.JwtService, error) {
	secretKey, err := generateSecretKey(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secret key: %v", err)
	}

	config, err := newJWTConfig(secretKey)
	if err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newJWTService(config.GetSecretKey())
}

func generateSecretKey(size int) (string, error) {
	secret := make([]byte, size)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(secret), nil
}
