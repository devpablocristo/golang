package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func BootstrapMicroService(server ports.GrpcServer, client ports.GrpcClient) (ports.GrpcService, error) {
	config := newConfigGrpcService(
		viper.GetString("MICRO_SERVICE_NAME"),
		server.Server(),
		client.Client(),
		viper.GetString("CONSUL_ADDRESS"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newGrpcService(config)
}
