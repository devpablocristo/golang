package sdkgomicro

import (
	"fmt"
	"sync"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
	"github.com/go-micro/plugins/v4/registry/consul"
)

var (
	instance  ports.GrpcService
	once      sync.Once
	initError error
)

type service struct {
	grpcService micro.Service
}

func newGrpcService(config ports.ConfigGrpcService) (ports.GrpcService, error) {
	once.Do(func() {
		if err := config.Validate(); err != nil {
			initError = fmt.Errorf("config validation error: %w", err)
			return
		}

		instance = &service{
			grpcService: setupGrpcService(config),
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcService(config ports.ConfigGrpcService) micro.Service {
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

func (s *service) Service() micro.Service {
	return s.grpcService
}

func setupRegistry(config ports.ConfigGrpcService) registry.Registry {
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.GetConsulAddress()}
	})
	return consulReg
}
