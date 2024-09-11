package jwtpkg

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
)

type jwtConfig struct {
	secretKey string
}

func newJWTConfig(secretKey string) (ports.JwtConfig, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("secret key cannot be empty")
	}
	return &jwtConfig{
		secretKey: secretKey,
	}, nil
}

func (config *jwtConfig) GetSecretKey() string {
	return config.secretKey
}

func (config *jwtConfig) Validate() error {
	if config.secretKey == "" {
		return fmt.Errorf("JWT secret key is not configured")
	}
	return nil
}
