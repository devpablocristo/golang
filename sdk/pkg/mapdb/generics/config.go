package mapdb

import (
	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

type mapDbConfig[T any] struct {
	prepopulate     bool
	prepopulateData []T
}

func NewMapDbConfig[T any](prepopulate bool, data []T) portspkg.MapDbConfig[T] {
	return &mapDbConfig[T]{
		prepopulate:     prepopulate,
		prepopulateData: data,
	}
}

func (c *mapDbConfig[T]) GetPrepopulate() bool {
	return c.prepopulate
}

func (c *mapDbConfig[T]) SetPrepopulate(prepopulate bool) {
	c.prepopulate = prepopulate
}

func (c *mapDbConfig[T]) GetPrepopulateData() []T {
	return c.prepopulateData
}

func (c *mapDbConfig[T]) SetPrepopulateData(data []T) {
	c.prepopulateData = data
}

func (c *mapDbConfig[T]) Validate() error {
	// Validar la configuración según sea necesario
	return nil
}
