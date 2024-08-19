package mapdb

import (
	"errors"
	"sync"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

var (
	instance portspkg.MapDbClient
	once     sync.Once
	errInit  error
)

type mapDbClient struct {
	db map[string]interface{}
}

// InitializeMapDbClient inicializa el cliente MapDb con la configuración dada.
func InitializeMapDbClient() error {
	once.Do(func() {
		client := &mapDbClient{}
		client.initialize()
	})
	return errInit
}

// Initialize inicializa la base de datos en memoria.
func (c *mapDbClient) initialize() {
	c.db = make(map[string]interface{})
}

// GetDB retorna la base de datos en memoria.
func (c *mapDbClient) GetDb() map[string]interface{} {
	return c.db
}

// GetMapDbInstance retorna la instancia inicializada de MapDbClient.
func GetMapDbInstance() (portspkg.MapDbClient, error) {
	if instance == nil {
		return nil, errors.New("mapdb client is not initialized")
	}
	return instance, nil
}
