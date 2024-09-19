package sdkgomicro

import (
	"fmt"

	grpcclient "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	grpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
	ginport "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type config struct {
	grpcClient      grpcclient.Client
	grpcServer      grpcserver.Server
	ginServer       ginport.Server
	grpcServiceName string
	webServiceName  string
	consulAddress   string
}

func newConfig(grpcClient grpcclient.Client, grpcServer grpcserver.Server, ginServer ginport.Server, grpcServiceName, webServiceName, consulAddress string) ports.Config {
	return &config{
		grpcClient:      grpcClient,
		grpcServer:      grpcServer,
		ginServer:       ginServer,
		grpcServiceName: grpcServiceName,
		webServiceName:  webServiceName,
		consulAddress:   consulAddress,
	}
}

func (config *config) GetGrpcServiceName() string {
	return config.grpcServiceName
}

func (config *config) GetWebServiceName() string {
	return config.webServiceName
}

func (config *config) GetConsulAddress() string {
	return config.consulAddress
}

func (config *config) Validate() error {
	if config.grpcServiceName == "" {
		return fmt.Errorf("missing grpc service name")
	}
	if config.webServiceName == "" {
		return fmt.Errorf("missing gin server name")
	}
	return nil
}
