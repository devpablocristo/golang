package gomicropkg

import (
	"fmt"

	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/events"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/store"
	syncm "go-micro.dev/v4/sync"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/web"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/portspkg"
)

// goMicroConfig representa la configuración necesaria para un servicio Go Micro.
type goMicroConfig struct {
	Name       string
	Version    string
	Address    string
	Registry   registry.Registry
	Logger     logger.Logger
	Auth       auth.Auth
	Broker     broker.Broker
	Client     client.Client
	Server     server.Server
	Store      store.Store
	Transport  transport.Transport
	WebService web.Service
	Config     config.Config
	Selector   selector.Selector
	Sync       syncm.Sync
	Events     events.Stream
}

// NewGoMicroConfig crea una nueva configuración de Go Micro con los valores obligatorios.
func NewGoMicroConfig(name, version, address string) portspkg.GoMicroConfig {
	return &goMicroConfig{
		Name:    name,
		Version: version,
		Address: address,
	}
}

// GetName devuelve el nombre del servicio.
func (config *goMicroConfig) GetName() string {
	return config.Name
}

// GetVersion devuelve la versión del servicio.
func (config *goMicroConfig) GetVersion() string {
	return config.Version
}

// GetAddress devuelve la dirección del servicio.
func (config *goMicroConfig) GetAddress() string {
	return config.Address
}

func (config *goMicroConfig) SetRegistry(reg registry.Registry) {
	config.Registry = reg
}

func (config *goMicroConfig) SetAuth(auth auth.Auth) {
	config.Auth = auth
}

func (config *goMicroConfig) SetBroker(broker broker.Broker) {
	config.Broker = broker
}

func (config *goMicroConfig) SetClient(client client.Client) {
	config.Client = client
}

func (config *goMicroConfig) SetLogger(logger logger.Logger) {
	config.Logger = logger
}

func (config *goMicroConfig) SetServer(server server.Server) {
	config.Server = server
}

func (config *goMicroConfig) SetStore(store store.Store) {
	config.Store = store
}

func (config *goMicroConfig) SetTransport(transport transport.Transport) {
	config.Transport = transport
}

func (config *goMicroConfig) SetWebService(webService web.Service) {
	config.WebService = webService
}

func (config *goMicroConfig) SetConfig(conf config.Config) {
	config.Config = conf
}

func (config *goMicroConfig) SetSelector(selector selector.Selector) {
	config.Selector = selector
}

func (config *goMicroConfig) SetSync(sync syncm.Sync) {
	config.Sync = sync
}

func (config *goMicroConfig) SetEvents(events events.Stream) {
	config.Events = events
}

// Validate valida que los valores obligatorios estén configurados.
func (config *goMicroConfig) Validate() error {
	if config.Name == "" {
		return fmt.Errorf("service name is not configured")
	}
	if config.Version == "" {
		return fmt.Errorf("service version is not configured")
	}
	if config.Address == "" {
		return fmt.Errorf("service address is not configured")
	}
	return nil
}
