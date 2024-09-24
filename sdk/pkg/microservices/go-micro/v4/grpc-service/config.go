package sdkgomicro

import (
	"fmt"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service/ports"
)

type config struct {
	serviceName   string
	server        server.Server
	client        client.Client
	consulAddress string
}

func newConfig(
	serviceName string,
	server server.Server,
	client client.Client,
	consulAddress string,
) ports.Config {
	return &config{
		serviceName:   serviceName,
		server:        server,
		client:        client,
		consulAddress: consulAddress,
	}
}

func (c *config) GetServiceName() string {
	return c.serviceName
}

func (c *config) GetServer() server.Server {
	return c.server
}

func (c *config) GetClient() client.Client {
	return c.client
}

func (c *config) GetConsulAddress() string {
	return c.consulAddress
}

func (c *config) Validate() error {
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
