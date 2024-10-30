package defs

import (
	"net/smtp"
	"time"
)

// Config define la interfaz que debe cumplir la configuración SMTP
type Config interface {
	GetSMTPServer() string
	GetAuth() smtp.Auth
	GetFrom() string
	Validate() error
}

type Service interface {
	SendVerificationEmail(to string, expiration time.Duration) error
}
