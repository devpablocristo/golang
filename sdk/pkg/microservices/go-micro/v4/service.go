package gomicropkg

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-micro/plugins/v4/registry/consul"
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
	rpcService micro.Service
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

		instance = &service{
			rpcService: setupRcpService(config),
			webService: setupWebService(config),
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

func setupRcpService(config ports.Config) micro.Service {
	rcpService := micro.NewService(
		micro.Name(config.GetRcpServiceName()),
		micro.Address(config.GetRcpServiceAddress()),
		micro.Registry(setupRegistry(config)),
	)
	return rcpService
}

func setupWebService(config ports.Config) web.Service {
	webService := web.NewService(
		web.Name(config.GetWebServiceName()),
		web.Address(config.GetWebServiceAddress()),
		web.Registry(setupRegistry(config)),
	)
	return webService
}

func setupRegistry(config ports.Config) registry.Registry {
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.GetConsulAddress()}
	})
	return consulReg
}
func setupLogger(config ports.Config) logger.Logger {
	loggerService := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithOutput(os.Stdout),
	)

	logger.DefaultLogger = logger.NewLogger(logger.WithLevel(logger.DebugLevel))

	return loggerService
}

func (s *service) StartRcpService() error {
	if s.rpcService != nil {
		err := s.rpcService.Run()
		if err != nil {
			return fmt.Errorf("failed to start rcp service: %w", err)
		}
		return nil
	}
	return fmt.Errorf("rpc service is not initialized")
}

func (s *service) StartWebService() error {
	if s.webService != nil {
		err := s.webService.Run()
		if err != nil {
			return fmt.Errorf("failed to start web service: %w", err)
		}
		return nil
	}
	return fmt.Errorf("web service is not initialized")
}

func (s *service) GetRcpService() micro.Service { return s.rpcService }
func (s *service) GetWebService() web.Service   { return s.webService }

// func (s *service) GetAuth() auth.Auth             { return config.Auth }
// func (s *service) GetBroker() broker.Broker       { return config.Broker }
// func (s *service) GetRegistry() registry.Registry { return config.Registry }
// func (s *service) GetLogger() logger.Logger       { return config.Logger }
// func (s *service) GetWebService() web.Service     { return config.WebService }

// func (s *service) SetService(service micro.Service)           { config.Service = service }
// func (s *service) SetRegistry(reg registry.Registry)          { config.Registry = reg }
// func (s *service) SetAuth(auth auth.Auth)                     { config.Auth = auth }
// func (s *service) SetBroker(broker broker.Broker)             { config.Broker = broker }
// func (s *service) SetClient(client client.Client)             { config.Client = client }
// func (s *service) SetLogger(logger logger.Logger)             { config.Logger = logger }
// func (s *service) SetServer(server server.Server)             { config.Server = server }
// func (s *service) SetStore(store store.Store)                 { config.Store = store }
// func (s *service) SetTransport(transport transport.Transport) { config.Transport = transport }
// func (s *service) SetWebService(webService web.Service)       { config.WebService = webService }
// func (s *service) SetConfig(conf configx.Config)              { config.Config = conf }
// func (s *service) SetSelector(selector selector.Selector)     { config.Selector = selector }
// func (s *service) SetSync(sync syncx.Sync)                    { config.Sync = sync }
// func (s *service) SetEvents(events events.Stream)             { config.Events = events }

// func (s *service) GetService() micro.Service {
// 	return s.service
// }

// func (s *service) GetGrpcClient() client.Client {
// 	return c.client
// }

// func (s *service) GetGrpcServer() server.Server {
// 	return c.server
// }

// func (s *service) GetAuth() auth.Auth {
// 	return c.auth
// }

// func (s *service) GetBroker() broker.Broker {
// 	return c.broker
// }

// func (s *service) GetConfig() configx.Config {
// 	return c.config
// }

// func (s *service) GetLogger() logger.Logger {
// 	return c.logger
// }

// func (s *service) GetRegistry() registry.Registry {
// 	return c.registry
// }

// func (s *service) GetSelector() selector.Selector {
// 	return c.selector
// }

// func (s *service) GetStore() store.Store {
// 	return c.store
// }
