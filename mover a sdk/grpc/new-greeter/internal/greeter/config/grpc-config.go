package config

import (
	"os"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
)

type ConfigGrpc struct {
	serverPort string
	serverHost string
}

func NewGrpcConfig() ctypes.ConfigGrpcPort {
	return &ConfigGrpc{
		serverPort: os.Getenv("GRPC_GREETER_SERVER_PORT"),
		serverHost: os.Getenv("GRPC_GREETER_SERVER_HOST"),
	}
}

func (c *ConfigGrpc) GetServerPort() string {
	return c.serverPort
}

func (c *ConfigGrpc) GetServerHost() string {
	return c.serverHost
}
