package sdkgomicro

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/web-server/ports"

)

type config struct {
	webRouter     interface{}
	webServerName string
	consulAddress string
	webServerHost string
	webServerPort int
}

func newConfig(
	webRouter interface{},
	webServerName string,
	consulAddress string,
	webServerHost string,
	webServerPort int,
) ports.Config{
	return &config{
		webRouter:     webRouter,
		webServerName: webServerName,
		consulAddress: consulAddress,
		webServerHost: webServerHost,
		webServerPort: webServerPort,
	}
}

func (c *config) GetRouter() interface{} {
	return c.webRouter
}

func (c *config) GetServerName() string {
	return c.webServerName
}

func (c *config) GetServerHost() string {
	return c.webServerHost
}

func (c *config) GetServerPort() int {
	return c.webServerPort
}

func (c *config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.webServerHost, c.webServerPort)
}

func (c *config) GetConsulAddress() string {
	return c.consulAddress
}

func (c *config) Validate() error {
	if c.webServerName == "" {
		return fmt.Errorf("missing web server name")
	}
	if c.webServerHost == "" {
		return fmt.Errorf("missing web server host")
	}
	if c.webServerPort == 0 {
		return fmt.Errorf("missing web server port")
	}
	if c.consulAddress == "" {
		return fmt.Errorf("missing consul address")
	}
	return nil
}
