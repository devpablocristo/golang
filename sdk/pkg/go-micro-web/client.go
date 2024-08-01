package gomicro

import (
	"fmt"
	"sync"

	"go-micro.dev/v4/web"
)

var (
	instance GoMicroClientPort
	once     sync.Once
	errInit  error
)

type GoMicroClientPort interface {
	Start() error
	Stop() error
	GetService() web.Service
}

type GoMicroClient struct {
	service web.Service
}

func InitializeGoMicroClient(config GoMicroConfig) error {
	once.Do(func() {
		ms := web.NewService(
			web.Name(config.Name),
			web.Version(config.Version),
			web.Registry(config.Registry),
			web.Address(config.Address),
		)

		if err := ms.Init(); err != nil {
			errInit = fmt.Errorf("failed to initialize go micro service: %w", err)
			return
		}
		instance = &GoMicroClient{service: ms}
	})
	return errInit
}

func GetGoMicroInstance() (GoMicroClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("go micro client is not initialized")
	}
	return instance, nil
}

func (c *GoMicroClient) Start() error {
	return c.service.Run()
}

func (c *GoMicroClient) Stop() error {
	return c.service.Stop()
}

func (c *GoMicroClient) GetService() web.Service {
	return c.service
}