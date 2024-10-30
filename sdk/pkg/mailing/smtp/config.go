package sdksmtp

import (
	"fmt"
	"net/smtp"

	defs "github.com/devpablocristo/golang/sdk/pkg/mailing/smtp/defs"
)

type config struct {
	smtpServer string
	auth       smtp.Auth
	from       string
}

func newConfig(smtpServer, from, username, password, port, identity string) defs.Config {
	auth := smtp.PlainAuth(identity, username, password, smtpServer)

	return &config{
		smtpServer: smtpServer + ":" + port,
		auth:       auth,
		from:       from,
	}
}

// GetSMTPServer devuelve la dirección del servidor SMTP con el puerto configurado
func (c *config) GetSMTPServer() string {
	return c.smtpServer
}

// GetAuth devuelve la autenticación SMTP configurada
func (c *config) GetAuth() smtp.Auth {
	return c.auth
}

// GetFrom devuelve la dirección de correo del remitente
func (c *config) GetFrom() string {
	return c.from
}

// Validate verifica que la configuración sea válida
func (c *config) Validate() error {
	if c.smtpServer == "" {
		return fmt.Errorf("SMTP server is not configured")
	}
	if c.auth == nil {
		return fmt.Errorf("SMTP auth is not configured")
	}
	if c.from == "" {
		return fmt.Errorf("SMTP from address is not configured")
	}
	return nil
}
