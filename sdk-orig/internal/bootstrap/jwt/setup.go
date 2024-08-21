package jwtsetup

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	jwtpkg "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/portspkg"
)

func NewJWTInstance() (portspkg.JwtClient, error) {
	secretKey, err := generateSecretKey(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secret key: %v", err)
	}

	config, err := jwtpkg.NewJWTConfig(secretKey)
	if err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := jwtpkg.InitializeJWTService(config.GetSecretKey()); err != nil {
		return nil, err
	}

	return jwtpkg.GetJWTInstance()
}

func generateSecretKey(size int) (string, error) {
	secret := make([]byte, size)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(secret), nil
}
