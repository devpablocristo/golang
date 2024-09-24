package sdkgomicro

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server/ports"
)

type config struct {
	grpcServerName string
	grpcServerHost string
	grpcServerPort int
}

func newConfig(grpcServerName string, grpcServerHost string, grpcServerPort int) ports.Config {
	return &config{
		grpcServerName: grpcServerName,
		grpcServerHost: grpcServerHost,
		grpcServerPort: grpcServerPort,
	}
}

func (c *config) GetGrpcServerName() string {
	return c.grpcServerName
}

func (c *config) GetGrpcServerHost() string {
	return c.grpcServerHost
}

func (c *config) GetGrpcServerPort() int {
	return c.grpcServerPort
}

func (c *config) Validate() error {
	if c.grpcServerName == "" {
		return fmt.Errorf("missing grpc service name")
	}
	if c.grpcServerHost == "" {
		return fmt.Errorf("missing grpc server host")
	}
	if c.grpcServerPort == 0 {
		return fmt.Errorf("missing grpc server port")
	}
	return nil
}
