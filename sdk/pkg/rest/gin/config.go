package ginpkg

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type config struct {
	routerPort string
}

func NewConfig(routerPort string) ports.Config {
	return &config{
		routerPort: routerPort,
	}
}

func (config *config) GetRouterPort() string {
	return config.routerPort
}

func (config *config) SetRouterPort(routerPort string) {
	config.routerPort = routerPort
}

func (config *config) Validate() error {
	if config.routerPort == "" {
		return fmt.Errorf("router port is not configured")
	}
	return nil
}
