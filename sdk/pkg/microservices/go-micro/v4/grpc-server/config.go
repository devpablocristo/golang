package sdkgomicro

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server/ports"
	"github.com/google/uuid"
)

type config struct {
	ServerName string
	ServerHost string
	ServerPort int
	ServerID   string
}

func newConfig(ServerName string, ServerHost string, ServerPort int) ports.Config {
	return &config{
		ServerName: ServerName,
		ServerHost: ServerHost,
		ServerPort: ServerPort,
		ServerID:   uuid.New().String(),
	}
}

func (c *config) GetServerName() string {
	return c.ServerName
}

func (c *config) GetServerHost() string {
	return c.ServerHost
}

func (c *config) GetServerPort() int {
	return c.ServerPort
}

func (c *config) GetServerID() string {
	return c.ServerID
}

func (c *config) Validate() error {
	if c.ServerName == "" {
		return fmt.Errorf("missing  service name")
	}
	if c.ServerHost == "" {
		return fmt.Errorf("missing  server host")
	}
	if c.ServerPort == 0 {
		return fmt.Errorf("missing  server port")
	}
	return nil
}
