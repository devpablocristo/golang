package smtp

import (
	"fmt"

	"github.com/spf13/viper"

	defs "github.com/devpablocristo/golang/sdk/pkg/mailing/smtp/defs"
)

func Bootstrap() (defs.Service, error) {
	config := newConfig(
		viper.GetString("SMTP_SERVER"),
		viper.GetString("SMTP_PORT"),
		viper.GetString("SMTP_FROM"),
		viper.GetString("SMTP_USERNAME"),
		viper.GetString("SMTP_PASSWORD"),
		viper.GetString("SMTP_IDENTITY"),
	)

	// Validar la configuración
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("Error en la configuración SMTP: %w", err)
	}

	// Crear el servicio SMTP con la configuración
	return newService(config)
}
