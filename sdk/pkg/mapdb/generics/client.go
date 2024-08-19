package mapdb

import (
	"context"
	"errors"
	"sync"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

type mapDbClient[T any] struct {
	db map[string]T
}

var (
	instance interface{}
	once     sync.Once
	errInit  error
)

func InitializeMapDbClient[T any](config portspkg.MapDbConfig[T]) error {
	once.Do(func() {
		client := &mapDbClient[T]{}
		errInit = client.Initialize(config)
		if errInit == nil {
			instance = client
		}
	})
	return errInit
}

func GetMapDbInstance[T any]() (*mapDbClient[T], error) {
	if instance == nil {
		return nil, errors.New("mapdb client is not initialized")
	}
	return instance.(*mapDbClient[T]), nil
}

func (c *mapDbClient[T]) Initialize(config portspkg.MapDbConfig[T]) error {
	c.db = make(map[string]T)

	if config.GetPrepopulate() {
		prepopulateData := config.GetPrepopulateData()
		c.prepopulate(prepopulateData)
	}

	return nil
}

func (c *mapDbClient[T]) prepopulate(data []T) {
	for _, item := range data {
		id := extractID(item)
		c.db[id] = item
	}
}

// Métodos Save, Get, Delete, List y Update implementados según la interfaz MapDbClient

func (c *mapDbClient[T]) Save(ctx context.Context, id string, entity T) error {
	c.db[id] = entity
	return nil
}

func (c *mapDbClient[T]) Get(ctx context.Context, id string) (T, error) {
	entity, exists := c.db[id]
	if !exists {
		var zeroValue T
		return zeroValue, errors.New("entity not found")
	}
	return entity, nil
}

func (c *mapDbClient[T]) Delete(ctx context.Context, id string) error {
	delete(c.db, id)
	return nil
}

func (c *mapDbClient[T]) List(ctx context.Context) ([]T, error) {
	var entities []T
	for _, entity := range c.db {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (c *mapDbClient[T]) Update(ctx context.Context, id string, entity T) error {
	if _, exists := c.db[id]; !exists {
		return errors.New("entity not found")
	}
	c.db[id] = entity
	return nil
}

// Función extractID genérica, que deberás adaptar a tu tipo T.
func extractID[T any](item T) string {
	// Implementa esto según la estructura de tus datos.
	return ""
}
