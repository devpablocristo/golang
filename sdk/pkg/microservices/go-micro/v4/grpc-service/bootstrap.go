package sdkgomicro

import (
	"github.com/spf13/viper"

	sdkclient "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client/ports"
	sdkserver "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server/ports"
	sdkservice "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service/ports"
)

func Bootstrap(server sdkserver.Server, client sdkclient.Client) (sdkservice.Service, error) {
	config := newConfig(
		server.GetServer(),
		client.GetClient(),
		viper.GetString("CONSUL_ADDRESS"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}