package mailing

import (
	"time"

	ports "github.com/devpablocristo/golang/sdk/sg/mailing/internal/core/ports"
)

type UseCases struct {
	smtpService ports.SmtpService
}

func NewUseCases(ss ports.SmtpService) ports.UseCases {
	return &UseCases{
		smtpService: ss,
	}
}

func (u *UseCases) SendVerificationEmail(email string, expiration time.Duration) error {
	return u.smtpService.SendVerificationEmail(email, expiration)
}
