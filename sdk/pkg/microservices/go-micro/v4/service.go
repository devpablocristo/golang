package gomicropkg

import (
	"fmt"
	"sync"

	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	configx "go-micro.dev/v4/config"
	"go-micro.dev/v4/events"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/store"
	syncx "go-micro.dev/v4/sync"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/web"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	service    micro.Service
	webService web.Service
	client     client.Client
	server     server.Server
	auth       auth.Auth
	broker     broker.Broker
	config     configx.Config
	logger     logger.Logger
	registry   registry.Registry
	store      store.Store
	transport  transport.Transport
	sync       syncx.Sync
	events     events.Stream
	selector   selector.Selector
}

func NewService(config ports.Config) (ports.Service, error) {
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

		instance = &service{
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

func GetInstance() (ports.Service, error) {
	if instance == nil {
		return nil, fmt.Errorf("go micro service is not initialized")
	}
	return instance, nil
}

func (c *service) Start() error {
	if c.webService != nil {
		return c.webService.Run()
	}
	return fmt.Errorf("web service is not initialized")
}

func (c *service) Stop() error {
	if c.webService != nil {
		return c.webService.Stop()
	}
	return fmt.Errorf("web service is not initialized")
}

func (c *service) GetService() micro.Service {
	return c.service
}

func (c *service) GetWebService() web.Service {
	return c.webService
}

func (c *service) GetGrpcClient() client.Client {
	return c.client
}

func (c *service) GetGrpcServer() server.Server {
	return c.server
}

func (c *service) GetAuth() auth.Auth {
	return c.auth
}

func (c *service) GetBroker() broker.Broker {
	return c.broker
}

func (c *service) GetConfig() configx.Config {
	return c.config
}

func (c *service) GetLogger() logger.Logger {
	return c.logger
}

func (c *service) GetRegistry() registry.Registry {
	return c.registry
}

func (c *service) GetSelector() selector.Selector {
	return c.selector
}

func (c *service) GetStore() store.Store {
	return c.store
}

func (c *service) GetTransport() transport.Transport {
	return c.transport
}

func (c *service) GetEvents() events.Stream {
	return c.events
}
