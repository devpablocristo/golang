package sdkgomicro

import (
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	grpcClient "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	grpcServer "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
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
	webServer   web.Service
}

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		setupLogger()

		instance = &service{
			grpcService: setupGrpcService(config),
			webServer:   setupWebServer(config),
		}

		err := instance.SetWebRouter(config.GetWebRouter())
		if err != nil {
			initError = fmt.Errorf("error setting web router: %w", err)
			return
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcService(config ports.Config) micro.Service {
	grpcSrv := grpcServer.NewServer(
		server.Address(fmt.Sprintf("%s:%d", config.GetGrpcServerHost(), config.GetGrpcServerPort())),
	)

	grpcClt := grpcClient.NewClient()

	grpcService := micro.NewService(
		micro.Name(config.GetGrpcServiceName()),
		micro.Server(grpcSrv),
		micro.Client(grpcClt),
		micro.Registry(setupRegistry(config)),
	)

	grpcService.Init()

	return grpcService
}

func setupWebServer(config ports.Config) web.Service {
	webServer := web.NewService(
		web.Name(config.GetWebServerName()),
		web.Address(config.GetWebServerAddress()),
		web.Registry(setupRegistry(config)),
	)

	return webServer
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

func (s *service) SetWebRouter(router interface{}) error {
	switch r := router.(type) {
	case *gin.Engine:
		s.webServer.Handle("/", r)
	default:
		return fmt.Errorf("unsupported router type")
	}
	return nil
}
