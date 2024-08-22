package pkgmapdb

import (
	"errors"
	"sync"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
)

var (
	instance ports.Service
	once     sync.Once
)

type service struct {
	db map[string]any
}

func NewService() ports.Service {
	once.Do(func() {
		instance = &service{
			db: make(map[string]any),
		}
	})
	return instance
}

func GetInstance() (ports.Service, error) {
	if instance == nil {
		return nil, errors.New("mapdb client is not initialized")
	}
	return instance, nil
}

func (c *service) GetDb() map[string]any {
	return c.db
}
