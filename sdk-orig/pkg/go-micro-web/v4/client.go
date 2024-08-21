package gomicro

import (
	"fmt"
	"sync"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/web"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/go-micro-web/v4/portspkg"
)

var (
	instance portspkg.GoMicroClient
	once     sync.Once
	errInit  error
)

type goMicroClient struct {
	service    micro.Service
	webService web.Service
	client     client.Client
	server     server.Server
}

// InitializeGoMicroClient inicializa una instancia única del cliente y servidor Go-Micro.
func InitializeGoMicroClient(config portspkg.GoMicroConfig) error {
	once.Do(func() {
		ms := micro.NewService(
			micro.Name(config.GetName()),
			micro.Version(config.GetVersion()),
			micro.Registry(config.GetRegistry()),
			micro.Address(config.GetAddress()),
		)

		ms.Init()

		instance = &goMicroClient{
			service: ms,
			client:  ms.Client(),
			server:  ms.Server(),
		}
	})
	return errInit
}

// GetGoMicroClientInstance devuelve la instancia del cliente Go-Micro.
func GetGoMicroClientInstance() (portspkg.GoMicroClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("go micro client is not initialized")
	}
	return instance, nil
}

// Start inicia el servicio Go-Micro.
func (c *goMicroClient) Start() error {
	return c.webService.Run()
}

// Stop detiene el servicio Go-Micro.
func (c *goMicroClient) Stop() error {
	return nil
}

// GetService devuelve la instancia del servicio Go-Micro.
func (c *goMicroClient) GetService() micro.Service {
	return c.service
}

// GetWebService devuelve la instancia del servicio web Go-Micro.
func (c *goMicroClient) GetWebService() web.Service {
	return c.webService
}

// GetGrpcClient devuelve la instancia del cliente gRPC.
func (c *goMicroClient) GetGrpcClient() client.Client {
	return c.client
}

// GetGrpcServer devuelve la instancia del servidor gRPC.
func (c *goMicroClient) GetGrpcServer() server.Server {
	return c.server
}
