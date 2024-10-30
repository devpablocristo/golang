package ports

import "time"

type SmtpService interface {
	SendVerificationEmail(email string, expiration time.Duration) error
}

type UseCases interface {
	SendVerificationEmail(email string, expiration time.Duration) error
}
