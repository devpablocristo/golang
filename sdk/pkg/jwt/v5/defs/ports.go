// defs/types.go

package defs

import "time"

// Config define la interfaz para la configuración del servicio JWT.
type Config interface {
	GetSecretKey() string
	Validate() error
}

// Service define la interfaz para el servicio JWT.
type Service interface {
	GenerateTokenForSubject(subject string, expiration time.Duration) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
}
