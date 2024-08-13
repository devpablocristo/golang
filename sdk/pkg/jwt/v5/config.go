package jwtpkg

import (
	"fmt"
)

// jwtConfig representa la configuración necesaria para manejar JWT.
type jwtConfig struct {
	secretKey string
}

// NewJWTConfig crea una nueva configuración de JWT con la clave secreta proporcionada.
func NewJWTConfig(secretKey string) (*jwtConfig, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("secret key cannot be empty")
	}
	return &jwtConfig{
		secretKey: secretKey,
	}, nil
}

// GetSecretKey devuelve la clave secreta de la configuración.
func (config *jwtConfig) GetSecretKey() string {
	return config.secretKey
}

// SetSecretKey establece la clave secreta de la configuración.
func (config *jwtConfig) SetSecretKey(secretKey string) error {
	if secretKey == "" {
		return fmt.Errorf("secret key cannot be empty")
	}
	config.secretKey = secretKey
	return nil
}

// Validate valida la configuración de JWT.
func (config *jwtConfig) Validate() error {
	if config.secretKey == "" {
		return fmt.Errorf("JWT secret key is not configured")
	}
	return nil
}
