package sdkgomicro

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
	grpcService micro.Service
	ginServer   web.Service
	client      client.Client
	server      server.Server
	auth        auth.Auth
	broker      broker.Broker
	config      configx.Config
	logger      logger.Logger
	registry    registry.Registry
	store       store.Store
	transport   transport.Transport
	sync        syncx.Sync
	events      events.Stream
	selector    selector.Selector
}

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		if err := config.Validate(); err != nil {
			initError = fmt.Errorf("config validation error: %w", err)
			return
		}
		setupLogger()

		instance = &service{
			grpcService: setupRcpService(config),
			ginServer:   setupGinServer(config),
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcService(config ports.Config) micro.Service {
	// Configuración del servidor gRPC
	grpcService := micro.NewService(
		micro.Name(config.GetGrpcServiceName()),
		micro.Address(config.GetGrpcServiceAddress()),
		micro.Server(grpcserver.NewServer()),
		micro.Client(grpcclient.NewClient()),
		micro.Registry(setupRegistry(config)),
		micro.
	)

	grpcService.Init()

	return grpcService
}

func setupginServer(config ports.Config) web.Service {
	ginServer := web.NewService(
		web.Name(config.GetginServerName()),
		web.Address(config.GetginServerAddress()),
		web.Registry(setupRegistry(config)),	
	)
	return ginServer
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

func (s *service) StartGrpcService() error {
	if s.grpcService != nil {
		err := s.grpcService.Run()
		if err != nil {
			return fmt.Errorf("failed to start rcp service: %w", err)
		}
		return nil
	}
	return fmt.Errorf("rpc service is not initialized")
}

func (s *service) StartginServer() error {
	if s.ginServer != nil {
		err := s.ginServer.Run()
		if err != nil {
			return fmt.Errorf("failed to start web service: %w", err)
		}
		return nil
	}
	return fmt.Errorf("web service is not initialized")
}

func (s *service) GetGrcpService() micro.Service { return s.grpcService }
func (s *service) GetginServer() web.Service     { return s.ginServer }
