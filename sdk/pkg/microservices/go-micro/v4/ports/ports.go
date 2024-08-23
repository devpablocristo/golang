package gomicropkgports

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

type Service interface {
	StartWebService() error
	//StopWebService() error
	StartRcpService() error
	GetRcpService() micro.Service
	GetWebService() web.Service
	// GetGrpcClient() client.Client
	// GetGrpcServer() server.Server
	// GetAuth() auth.Auth
	// GetBroker() broker.Broker
	// GetConfig() config.Config
	// GetLogger() logger.Logger
	// GetRegistry() registry.Registry
	// GetSelector() selector.Selector
	// GetStore() store.Store
	// GetTransport() transport.Transport
	// GetEvents() events.Stream
}

type Config interface {
	// GetService() micro.Service
	// GetAuth() auth.Auth
	// GetBroker() broker.Broker
	// GetRegistry() registry.Registry
	// GetLogger() logger.Logger
	// GetWebService() web.Service
	// SetService(micro.Service)
	// SetRegistry(registry.Registry)
	// SetAuth(auth.Auth)
	// SetBroker(broker.Broker)
	// SetClient(client.Client)
	// SetLogger(logger.Logger)
	// SetServer(server.Server)
	// SetStore(store.Store)
	// SetTransport(transport.Transport)
	// SetWebService(web.Service)
	// SetConfig(config.Config)
	// SetSelector(selector.Selector)
	// SetSync(syncx.Sync)
	// SetEvents(events.Stream)
	GetRcpServiceName() string
	GetWebServiceName() string
	GetRcpServiceAddress() string
	GetWebServiceAddress() string
	GetConsulAddress() string
	Validate() error
}
