package mailgtw

import (
	"errors"

	"github.com/spf13/viper"
)

func getSecrets() (map[string]string, error) {
	// Crear un mapa para almacenar los secrets
	secrets := make(map[string]string)

	// Cargar los secrets cuando sea necesario
	afipSecret := viper.GetString("AFIP_CLIENT_SECRET")
	miArgSecret := viper.GetString("MIARG_CLIENT_SECRET")

	// Si los secrets están vacíos, retornamos error (por si acaso)
	if afipSecret == "" {
		return nil, errors.New("AFIP_CLIENT_SECRET is missing")
	}
	if miArgSecret == "" {
		return nil, errors.New("MIARG_CLIENT_SECRET is missing")
	}

	// Guardar los secretos en el mapa
	secrets["afip"] = afipSecret
	secrets["miarg"] = miArgSecret

	return secrets, nil
}
