package smtp

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"sync"
	"time"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	sdkjwtdefs "github.com/devpablocristo/golang/sdk/pkg/jwt/v5/defs"
	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"

	defs "github.com/devpablocristo/golang/sdk/pkg/mailing/smtp/defs"
)

var (
	instance defs.Service
	once     sync.Once
	initErr  error
)

// service representa el servicio SMTP que envía correos
type service struct {
	config     defs.Config
	jwtService sdkjwtdefs.Service
}

// newService crea una nueva instancia del servicio SMTP usando la configuración proporcionada
func newService(config defs.Config) (defs.Service, error) {
	once.Do(func() {

		js, err := sdkjwt.Bootstrap()
		if err != nil {
			initErr = fmt.Errorf("jwt bootstrap error: %w", err)
		}

		instance = &service{
			config:     config,
			jwtService: js,
		}
	})

	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}

func (s *service) SendVerificationEmail(to string, expiration time.Duration) error {
	// Validate email format
	if err := sdktools.ValidateEmail(to); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}

	// Generate a unique verification token
	token, err := s.jwtService.GenerateTokenForSubject(to, expiration)
	if err != nil {
		return fmt.Errorf("failed to generate token: %w", err)
	}
	verificationLink := fmt.Sprintf("https://your-domain.com/verify?token=%s", token)

	// Prepare the email subject and body
	subject := "Verify your email address"
	body := fmt.Sprintf("Click the following link to verify your email: %s", verificationLink)

	// Prepare the email message in the correct format
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))

	// Send the email using TLS
	serverAddr := s.config.GetSMTPServer()
	conn, err := tls.Dial("tcp", serverAddr, &tls.Config{
		InsecureSkipVerify: true, // In production, consider disabling this or use valid certificates
	})
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	// Create SMTP client
	client, err := smtp.NewClient(conn, s.config.GetSMTPServer())
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// SMTP authentication
	if err := client.Auth(s.config.GetAuth()); err != nil {
		return fmt.Errorf("failed to authenticate with SMTP server: %w", err)
	}

	// Set the sender and recipient
	if err := client.Mail(s.config.GetFrom()); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Write the email message
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get SMTP data writer: %w", err)
	}
	if _, err := w.Write(msg); err != nil {
		return fmt.Errorf("failed to write email message: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close email message writer: %w", err)
	}

	fmt.Printf("Verification email sent to %s\n", to)
	return nil
}
