package sdkgomicro

import (
	"os"
	"sync"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	grpcService micro.Service
}

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		setupLogger()

		instance = &service{
			grpcService: setupGrpcService(config),
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcService(config ports.Config) micro.Service {
	service := micro.NewService(
		micro.Name(config.GetServiceName()),
		micro.Server(config.GetServer()),
		micro.Client(config.GetClient()),
		micro.Registry(setupRegistry(config)),
	)

	service.Init()

	return service
}

func (s *service) Run() error {
	return s.grpcService.Run()
}

func (s *service) GetService() micro.Service {
	return s.grpcService
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
