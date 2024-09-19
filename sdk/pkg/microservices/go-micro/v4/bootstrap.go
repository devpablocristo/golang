package sdkgomicro

import (
	"github.com/spf13/viper"

	grpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	grpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
	ginport "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

// NOTE: rest (gin, std, chi), rpc(grpc, thrift), messaging(rabbitmq, kafka), websocket, graph ql
// TODO: Agregar rabbitmq y websocket
func Bootstrap(grpcClient grpcclient.Client, grpcServer grpcserver.Server, ginServer ginport.Server) (ports.Service, error) {
	config := newConfig(
		grpcClient,
		grpcServer,
		ginServer,
		viper.GetString("GRPC_SERVICE_NAME"),
		viper.GetString("WEB_SERVER_NAME"),
		viper.GetString("CONSUL_ADDRESS"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
