package mailconn

import (
	"log"
	"time"

	sdksmtp "github.com/devpablocristo/golang/sdk/pkg/mailing/smtp"
	sdlsmtpdefs "github.com/devpablocristo/golang/sdk/pkg/mailing/smtp/defs"

	ports "github.com/devpablocristo/golang/sdk/sg/mailing/internal/core/ports"
)

type SmtpService struct {
	smtpService sdlsmtpdefs.Service
}

func NewSmtpService() (ports.SmtpService, error) {
	// Inicializar el servicio de SMTP
	smtpService, err := sdksmtp.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize SMTP service: %v", err)
	}

	return &SmtpService{
		smtpService: smtpService,
	}, nil
}

// Implementación de SendVerificationEmail
func (ss *SmtpService) SendVerificationEmail(email string, expiration time.Duration) error {
	// Llamada al servicio SMTP para enviar el correo
	return ss.smtpService.SendVerificationEmail(email, expiration)
}
