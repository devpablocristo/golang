package mapdb

import (
	"errors"
	"sync"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

type mapDbClient struct {
	db map[string]interface{}
}

var (
	instance portspkg.MapDbConfig
	once     sync.Once
	errInit  error
)

// InitializeMapDbClient inicializa el cliente MapDb con la configuración dada.
func InitializeMapDbClient(config portspkg.MapDbConfig) error {
	once.Do(func() {
		client := &mapDbClient{}
		errInit = client.initialize(config)
		if errInit == nil {
			instance = client
		}
	})
	return errInit
}

// Initialize inicializa la base de datos en memoria.
func (c *mapDbClient) initialize(config portspkg.MapDbConfig) error {
	c.db = make(map[string]interface{})
	return nil // No hay validación en este caso, pero podrías agregarla si es necesario.
}

// GetMapDbInstance retorna la instancia inicializada de MapDbClient.
func GetMapDbInstance() (portspkg.MapDbClient, error) {
	if instance == nil {
		return nil, errors.New("mapdb client is not initialized")
	}
	return instance, nil
}
