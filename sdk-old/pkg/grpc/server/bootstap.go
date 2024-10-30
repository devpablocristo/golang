package sdkgrpcserver

import (
	"github.com/spf13/viper"

	defs "github.com/devpablocristo/golang/sdk/pkg/grpc/server/defs"
)

// Bootstrap inicializa y devuelve una instancia de servidor gRPC
func Bootstrap() (defs.Server, error) {
	config := newConfig(
		"", // viper.GetString("GRPC_SERVER_HOST"), // si es necesario
		viper.GetInt("GRPC_SERVER_PORT"),
		nil, // Configuración TLS, si es necesario
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newServer(config)
}
