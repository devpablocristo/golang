package sdkgomicro

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

type config struct {
	grpcServiceName    string
	grpcServiceAddress string
	ginServerName      string
	ginServerAddress   string
	consulAddress      string
}

func newConfig(grpcN, grpcA, ginN, ginA, cA string) ports.Config {
	return &config{
		grpcServiceName:    grpcN,
		grpcServiceAddress: grpcA,
		ginServerName:      ginN,
		ginServerAddress:   ginA,
		consulAddress:      cA,
	}
}

func (config *config) GetGrpcServiceName() string {
	return config.grpcServiceName
}

func (config *config) GetGinServerName() string {
	return config.ginServerName
}

func (config *config) GetGrpcServiceAddress() string {
	return config.grpcServiceAddress
}

func (config *config) GetGinServerAddress() string {
	return config.ginServerAddress
}

func (config *config) GetConsulAddress() string {
	return config.consulAddress
}

func (config *config) Validate() error {
	if config.grpcServiceName == "" && config.ginServerName == "" {
		return fmt.Errorf("missing service name: web or/and rpc")
	}
	return nil
}
