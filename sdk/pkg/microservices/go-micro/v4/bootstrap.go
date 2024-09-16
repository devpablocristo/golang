package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func Bootstrap(ginClient ginport.Server, grpcClient grpcclient.Client, grpcServer grpcserver.Server) (ports.Service, error) {

	
	

	config := newConfig(
		viper.GetString("GOMICRO_RPC_SERVICE_NAME"),
		":"+viper.GetString("GOMICRO_RPC_SERVICE_ADDRESS"),
		viper.GetString("GOMICRO_WEB_SERVICE_NAME"),
		":"+viper.GetString("GOMICRO_WEB_SERVICE_ADDRESS"),
		viper.GetString("CONSUL_ADDRESS"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
