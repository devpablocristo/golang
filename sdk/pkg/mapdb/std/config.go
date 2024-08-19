package mapdb

import (
	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

type mapDbConfig struct{}

func NewMapDbConfig() portspkg.MapDbConfig {
	return &mapDbConfig{}
}
