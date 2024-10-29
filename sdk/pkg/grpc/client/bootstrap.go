package sdkcgrpcclient

import (
	"github.com/spf13/viper"

	defs "github.com/devpablocristo/golang/sdk/pkg/grpc/client/defs"
)

func Bootstrap() (defs.Client, error) {
	config := newConfig(
		viper.GetString("GRPC_SERVER_HOST"),
		viper.GetInt("GRPC_SERVER_PORT"),
		nil, // Configuración TLS, si es necesario
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
