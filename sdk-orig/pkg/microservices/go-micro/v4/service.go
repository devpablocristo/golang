package gomicropkg

import (
	"fmt"
	"sync"

	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/events"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/store"
	syncm "go-micro.dev/v4/sync"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/web"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/portspkg"
)

var (
	instance  portspkg.GoMicroService
	once      sync.Once
	initError error
)

type goMicroService struct {
	service    micro.Service
	webService web.Service
	client     client.Client
	server     server.Server
	auth       auth.Auth
	broker     broker.Broker
	config     config.Config
	logger     logger.Logger
	registry   registry.Registry
	store      store.Store
	transport  transport.Transport
	sync       syncm.Sync
	events     events.Stream
	selector   selector.Selector
}

// NewGoMicroService crea una nueva instancia del servicio Go Micro con la configuración proporcionada.
func NewGoMicroService(config portspkg.GoMicroConfig) (portspkg.GoMicroService, error) {
	once.Do(func() {
		if err := config.Validate(); err != nil {
			initError = fmt.Errorf("config validation error: %w", err)
			return
		}

		ms := micro.NewService(
			micro.Name(config.GetName()),
			micro.Version(config.GetVersion()),
			micro.Address(config.GetAddress()),
		)

		ms.Init()

		if ms.Options().Registry == nil || ms.Client() == nil || ms.Server() == nil {
			initError = fmt.Errorf("failed to initialize Go-Micro service")
			return
		}

		instance = &goMicroService{
			service:   ms,
			client:    ms.Client(),
			server:    ms.Server(),
			auth:      ms.Options().Auth,
			broker:    ms.Options().Broker,
			config:    ms.Options().Config,
			logger:    ms.Options().Logger,
			registry:  ms.Options().Registry,
			store:     ms.Options().Store,
			transport: ms.Options().Transport,
			// Estos servicios adicionales deben configurarse según sea necesario
			sync:     nil,
			events:   nil,
			selector: nil,
		}
	})
	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

// GetGoMicroInstance devuelve la instancia del servicio Go-Micro.
func GetGoMicroInstance() (portspkg.GoMicroService, error) {
	if instance == nil {
		return nil, fmt.Errorf("go micro service is not initialized")
	}
	return instance, nil
}

// Implementación de los métodos de la interfaz GoMicroService

func (c *goMicroService) Start() error {
	if c.webService != nil {
		return c.webService.Run()
	}
	return fmt.Errorf("web service is not initialized")
}

func (c *goMicroService) Stop() error {
	if c.webService != nil {
		return c.webService.Stop()
	}
	return fmt.Errorf("web service is not initialized")
}

func (c *goMicroService) GetService() micro.Service {
	return c.service
}

func (c *goMicroService) GetWebService() web.Service {
	return c.webService
}

func (c *goMicroService) GetGrpcClient() client.Client {
	return c.client
}

func (c *goMicroService) GetGrpcServer() server.Server {
	return c.server
}

func (c *goMicroService) GetAuth() auth.Auth {
	return c.auth
}

func (c *goMicroService) GetBroker() broker.Broker {
	return c.broker
}

func (c *goMicroService) GetConfig() config.Config {
	return c.config
}

func (c *goMicroService) GetLogger() logger.Logger {
	return c.logger
}

func (c *goMicroService) GetRegistry() registry.Registry {
	return c.registry
}

func (c *goMicroService) GetSelector() selector.Selector {
	return c.selector
}

func (c *goMicroService) GetStore() store.Store {
	return c.store
}

func (c *goMicroService) GetTransport() transport.Transport {
	return c.transport
}

func (c *goMicroService) GetEvents() events.Stream {
	return c.events
}
