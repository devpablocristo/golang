package sdkjwt

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/defs"
)

type config struct {
	secretKey string
}

func newConfig(secretKey string) defs.Config {
	return &config{
		secretKey: secretKey,
	}
}

func (c *config) GetSecretKey() string {
	return c.secretKey
}

func (c *config) Validate() error {
	if c.secretKey == "" {
		return fmt.Errorf("JWT secret key is not configured")
	}
	return nil
}
