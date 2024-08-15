package gomicro

import (
	"fmt"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4/portspkg"
	"go-micro.dev/v4/registry"
)

type goMicroConfig struct {
	name     string
	version  string
	address  string
	registry registry.Registry
}

func NewGoMicroConfig(name, version, address string, reg registry.Registry) portspkg.GoMicroConfig {
	return &goMicroConfig{
		name:     name,
		version:  version,
		address:  address,
		registry: reg,
	}
}

func (c *goMicroConfig) GetName() string {
	return c.name
}

func (c *goMicroConfig) SetName(name string) {
	c.name = name
}

func (c *goMicroConfig) GetVersion() string {
	return c.version
}

func (c *goMicroConfig) SetVersion(version string) {
	c.version = version
}

func (c *goMicroConfig) GetAddress() string {
	return c.address
}

func (c *goMicroConfig) SetAddress(address string) {
	c.address = address
}

func (c *goMicroConfig) GetRegistry() any {
	return c.registry
}

func (c *goMicroConfig) SetRegistry(reg any) {
	c.registry = reg.(registry.Registry)
}

func (c *goMicroConfig) Validate() error {
	if c.name == "" {
		return fmt.Errorf("service name is not configured")
	}
	if c.version == "" {
		return fmt.Errorf("service version is not configured")
	}
	if c.address == "" {
		return fmt.Errorf("service address is not configured")
	}
	if c.registry == nil {
		return fmt.Errorf("service registry is not configured")
	}
	return nil
}
