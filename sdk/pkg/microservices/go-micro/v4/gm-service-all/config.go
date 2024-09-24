package sdkgomicro

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

type config struct {
	webRouter       interface{}
	grpcServerPort  int
	grpcServiceName string
	webServerName   string
	webServerHost   string
	webServerPort   int
	consulAddress   string
	grpcServerHost  string
}

func newConfig(
	webRouter interface{},
	grpcServerPort int,
	grpcServiceName string,
	webServerName string,
	consulAddress string,
	grpcServerHost string,
	webServerHost string,
	webServerPort int,
) ports.Config {
	return &config{
		webRouter:       webRouter,
		grpcServerPort:  grpcServerPort,
		grpcServiceName: grpcServiceName,
		webServerName:   webServerName,
		webServerHost:   webServerHost,
		webServerPort:   webServerPort,
		consulAddress:   consulAddress,
		grpcServerHost:  grpcServerHost,
	}
}

func (c *config) GetWebRouter() interface{} {
	return c.webRouter
}

func (c *config) GetGrpcServerPort() int {
	return c.grpcServerPort
}

func (c *config) GetGrpcServiceName() string {
	return c.grpcServiceName
}

func (c *config) GetWebServerName() string {
	return c.webServerName
}

func (c *config) GetWebServerHost() string {
	return c.webServerHost
}

func (c *config) GetWebServerPort() int {
	return c.webServerPort
}

func (c *config) GetWebServerAddress() string {
	return fmt.Sprintf("%s:%d", c.webServerHost, c.webServerPort)
}

func (c *config) GetConsulAddress() string {
	return c.consulAddress
}

func (c *config) GetGrpcServerHost() string {
	return c.grpcServerHost
}

func (c *config) Validate() error {
	if c.grpcServiceName == "" {
		return fmt.Errorf("missing grpc service name")
	}
	if c.grpcServerHost == "" {
		return fmt.Errorf("missing grpc server host")
	}
	if c.grpcServerPort == 0 {
		return fmt.Errorf("missing grpc server port")
	}
	if c.consulAddress == "" {
		return fmt.Errorf("missing consul address")
	}
	if c.webServerName == "" {
		return fmt.Errorf("missing web server name")
	}
	if c.webServerHost == "" {
		return fmt.Errorf("missing web server host")
	}
	if c.webServerPort == 0 {
		return fmt.Errorf("missing web server port")
	}
	return nil
}
