package sdkgin

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type config struct {
	routerPort   string
	apiVersion   string
	jwtSecretKey string
}

func newConfig(routerPort, apiVersion, jwtSecretKey string) ports.Config {
	return &config{
		routerPort:   routerPort,
		apiVersion:   apiVersion,
		jwtSecretKey: jwtSecretKey,
	}
}

func (c *config) GetRouterPort() string {
	return c.routerPort
}

func (c *config) SetRouterPort(routerPort string) {
	c.routerPort = routerPort
}

func (c *config) GetAPIVersion() string {
	return c.apiVersion
}

func (c *config) SetAPIVersion(apiVersion string) {
	c.apiVersion = apiVersion
}

func (c *config) GetJWTSecretKey() string {
	return c.jwtSecretKey
}

func (c *config) SetJWTSecretKey(jwtSecretKey string) {
	c.jwtSecretKey = jwtSecretKey
}

func (c *config) Validate() error {
	if c.routerPort == "" {
		return fmt.Errorf("router port is not configured")
	}
	if c.apiVersion == "" {
		return fmt.Errorf("API version is not configured")
	}
	if c.jwtSecretKey == "" {
		return fmt.Errorf("JWT secret key is not configured")
	}
	return nil
}
