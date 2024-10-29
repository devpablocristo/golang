package sdktools

import (
	"fmt"
	"regexp"
)

// ValidateEmail verifica si el formato del correo electrónico es válido
func ValidateEmail(email string) error {
	// Expresión regular para validar un correo electrónico
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email address: %s", email)
	}
	return nil
}
