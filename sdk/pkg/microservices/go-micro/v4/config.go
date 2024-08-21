package gomicropkg

import (
	"fmt"

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

// config representa la configuración necesaria para un servicio Go Micro.
type config struct {
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
	Config     configx.Config
	Selector   selector.Selector
	Sync       syncx.Sync
	Events     events.Stream
}

func NewConfig(name, version, address string) ports.Config {
	return &config{
		Name:    name,
		Version: version,
		Address: address,
	}
}

// GetName devuelve el nombre del servicio.
func (config *config) GetName() string {
	return config.Name
}

// GetVersion devuelve la versión del servicio.
func (config *config) GetVersion() string {
	return config.Version
}

// GetAddress devuelve la dirección del servicio.
func (config *config) GetAddress() string {
	return config.Address
}

func (config *config) SetRegistry(reg registry.Registry) {
	config.Registry = reg
}

func (config *config) SetAuth(auth auth.Auth) {
	config.Auth = auth
}

func (config *config) SetBroker(broker broker.Broker) {
	config.Broker = broker
}

func (config *config) SetClient(client client.Client) {
	config.Client = client
}

func (config *config) SetLogger(logger logger.Logger) {
	config.Logger = logger
}

func (config *config) SetServer(server server.Server) {
	config.Server = server
}

func (config *config) SetStore(store store.Store) {
	config.Store = store
}

func (config *config) SetTransport(transport transport.Transport) {
	config.Transport = transport
}

func (config *config) SetWebService(webService web.Service) {
	config.WebService = webService
}

func (config *config) SetConfig(conf configx.Config) {
	config.Config = conf
}

func (config *config) SetSelector(selector selector.Selector) {
	config.Selector = selector
}

func (config *config) SetSync(sync syncx.Sync) {
	config.Sync = sync
}

func (config *config) SetEvents(events events.Stream) {
	config.Events = events
}

// Validate valida que los valores obligatorios estén configurados.
func (config *config) Validate() error {
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
