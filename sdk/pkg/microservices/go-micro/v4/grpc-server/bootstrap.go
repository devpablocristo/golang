package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func BootstrapGrpcServer() (ports.GrpcServer, error) {
	config := newConfigGrpcServer(
		viper.GetString("GRPC_SERVICE_NAME"),
		viper.GetString("GRPC_SERVER_HOST"),
		viper.GetInt("GRPC_SERVER_PORT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newGrpcServer(config)
}
