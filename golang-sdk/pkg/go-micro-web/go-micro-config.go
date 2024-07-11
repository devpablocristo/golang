package gomicro

import (
	"go-micro.dev/v4/registry"
)

type GoMicroConfig struct {
	Name     string
	Version  string
	Address  string
	Registry registry.Registry
}
