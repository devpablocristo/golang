package sdkgomicro

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

type configGrpcServer struct {
	grpcServerName string
	grpcServerHost string
	grpcServerPort int
}

func newConfigGrpcServer(grpcServerName string, grpcServerHost string, grpcServerPort int) ports.ConfigGrpcServer {
	return &configGrpcServer{
		grpcServerName: grpcServerName,
		grpcServerHost: grpcServerHost,
		grpcServerPort: grpcServerPort,
	}
}

func (c *configGrpcServer) GetGrpcServerName() string {
	return c.grpcServerName
}

func (c *configGrpcServer) GetGrpcServerHost() string {
	return c.grpcServerHost
}

func (c *configGrpcServer) GetGrpcServerPort() int {
	return c.grpcServerPort
}

func (c *configGrpcServer) Validate() error {
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
