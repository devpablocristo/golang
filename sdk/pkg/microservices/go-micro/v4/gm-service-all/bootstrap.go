package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func Bootstrap(webRouter interface{}) (ports.Service, error) {
	config := newConfig(
		webRouter,
		viper.GetInt("GRPC_SERVER_PORT"),     // grpcServerPort
		viper.GetString("GRPC_SERVICE_NAME"), // grpcServiceName
		viper.GetString("WEB_SERVER_NAME"),   // webServerName
		viper.GetString("CONSUL_ADDRESS"),    // consulAddress
		viper.GetString("GRPC_SERVER_HOST"),  // grpcServerHost
		viper.GetString("WEB_SERVER_HOST"),   // webServerHost
		viper.GetInt("WEB_SERVER_PORT"),      // webServerPort
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
