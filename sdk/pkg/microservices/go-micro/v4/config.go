package pkggomicro

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

type config struct {
	rpcServiceName    string
	rpcServiceAddress string
	webServiceName    string
	webServiceAddress string
	consulAddress     string
}

func newConfig(rpcSN, rpcSA, webSN, webSA, cA string) ports.Config {
	return &config{
		rpcServiceName:    rpcSN,
		rpcServiceAddress: rpcSA,
		webServiceName:    webSN,
		webServiceAddress: webSA,
		consulAddress:     cA,
	}
}

func (config *config) GetRcpServiceName() string {
	return config.rpcServiceName
}

func (config *config) GetWebServiceName() string {
	return config.webServiceName
}

func (config *config) GetRcpServiceAddress() string {
	return config.rpcServiceAddress
}

func (config *config) GetWebServiceAddress() string {
	return config.webServiceAddress
}

func (config *config) GetConsulAddress() string {
	return config.consulAddress
}

func (config *config) Validate() error {
	if config.rpcServiceName == "" && config.webServiceName == "" {
		return fmt.Errorf("missing service name: web or/and rpc")
	}
	return nil
}
