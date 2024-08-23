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

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		if err := config.Validate(); err != nil {
			initError = fmt.Errorf("config validation error: %w", err)
			return
		}
		setupLogger()
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

func setupLogger() {
	logger.DefaultLogger = logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithOutput(os.Stdout),
	)
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
