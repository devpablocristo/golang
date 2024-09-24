package sdkgomicro

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	webServer web.Service
}

func newWebServer(config ports.ConfigWebServer) (ports.WebServer, error) {
	once.Do(func() {
		instance = &service{
			webServer: setupWebServer(config),
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

func setupWebServer(config ports.ConfigWebServer) web.Service {
	webServer := web.NewService(
		web.Name(config.GetWebServerName()),
		web.Address(config.GetWebServerAddress()),
		web.Registry(setupRegistry(config)),
	)

	return webServer
}

func setupRegistry(config ports.ConfigWebServer) registry.Registry {
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.GetConsulAddress()}
	})
	return consulReg
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
