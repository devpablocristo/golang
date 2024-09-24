package sdkgomicro

import (
	"fmt"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

type configGrpcService struct {
	serviceName   string
	server        server.Server
	client        client.Client
	consulAddress string
}

func newConfigGrpcService(
	serviceName string,
	server server.Server,
	client client.Client,
	consulAddress string,
) ports.ConfigGrpcService {
	return &configGrpcService{
		serviceName:   serviceName,
		server:        server,
		client:        client,
		consulAddress: consulAddress,
	}
}

func (c *configGrpcService) GetServiceName() string {
	return c.serviceName
}

func (c *configGrpcService) GetServer() server.Server {
	return c.server
}

func (c *configGrpcService) GetClient() client.Client {
	return c.client
}

func (c *configGrpcService) GetConsulAddress() string {
	return c.consulAddress
}

func (c *configGrpcService) Validate() error {
	if c.serviceName == "" {
		return fmt.Errorf("missing service name")
	}
	if c.server == nil {
		return fmt.Errorf("missing server")
	}
	if c.client == nil {
		return fmt.Errorf("missing client")
	}
	return nil
}
