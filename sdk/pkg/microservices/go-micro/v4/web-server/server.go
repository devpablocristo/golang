package sdkgomicro

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/web-server/ports"
)

var (
	instance  ports.Server
	once      sync.Once
	initError error
)

type service struct {
	server web.Service
}

func newServer(config ports.Config) (ports.Server, error) {
	once.Do(func() {
		instance = &service{
			server: setupWebServer(config),
		}

		err := instance.SetWebRouter(config.GetRouter())
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

func setupWebServer(config ports.Config) web.Service {
	Server := web.NewService(
		web.Name(config.GetServerName()),
		web.Address(config.GetServerAddress()),
		web.Registry(setupRegistry(config)),
	)

	return Server
}

func setupRegistry(config ports.Config) registry.Registry {
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.GetConsulAddress()}
	})
	return consulReg
}

func (s *service) SetWebRouter(router interface{}) error {
	switch r := router.(type) {
	case *gin.Engine:
		s.server.Handle("/", r)
	default:
		return fmt.Errorf("unsupported router type")
	}
	return nil
}

func (s *service) Run() error {
	return s.server.Run()
}
