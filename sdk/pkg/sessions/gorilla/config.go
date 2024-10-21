package sdksession

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla/ports"
)

type config struct {
	secretKey string
}

// newConfig crea una nueva configuración para Gorilla Sessions
func newConfig(secretKey string) ports.Config {
	return &config{
		secretKey: secretKey,
	}
}

// SecretKey retorna la clave secreta para encriptar las cookies de sesión
func (c *config) GetSecretKey() string {
	return c.secretKey
}

func (c *config) Validate() error {
	if c.secretKey == "" {
		return fmt.Errorf("GORILLA_SESSION_SECRET_KEY is required")
	}
	return nil
}
