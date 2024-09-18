package sdkgomicro

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"

	grpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	grpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
	ginport "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type config struct {
	grpcClient    grpcclient.Client
	grpcServer    grpcserver.Server
	ginClient     ginport.Server
	consulAddress string
}

func newConfig(grpcClient grpcclient.Client, grpcServer grpcserver.Server, ginServer ginport.Server, consulAddress string) ports.Config {
	return &config{
		grpcClient:    grpcClient,
		grpcServer:    grpcServer,
		ginClient:     ginServer,
		consulAddress: consulAddress,
	}
}

func (config *config) GetConsulAddress() string {
	return config.consulAddress
}

func (config *config) Validate() error {

	return nil
}
