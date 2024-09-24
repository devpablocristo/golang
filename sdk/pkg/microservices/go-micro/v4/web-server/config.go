package sdkgomicro

import (
	"fmt"
)

type config struct {
	webRouter     interface{}
	webServerName string
	consulAddress string
	webServerHost string
	webServerPort int
}

func newConfigWebServer(
	webRouter interface{},
	webServerName string,
	consulAddress string,
	webServerHost string,
	webServerPort int,
) *config {
	return &config{
		webRouter:     webRouter,
		webServerName: webServerName,
		consulAddress: consulAddress,
		webServerHost: webServerHost,
		webServerPort: webServerPort,
	}
}

func (c *config) GetWebRouter() interface{} {
	return c.webRouter
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
