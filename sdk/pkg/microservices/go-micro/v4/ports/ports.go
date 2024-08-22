package gomicropkgports

import (
	"go-micro.dev/v4"
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
	syncx "go-micro.dev/v4/sync"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/web"
)

type Service interface {
	Start() error
	Stop() error
	GetService() micro.Service
	GetWebService() web.Service
	GetGrpcClient() client.Client
	GetGrpcServer() server.Server
	GetAuth() auth.Auth
	GetBroker() broker.Broker
	GetConfig() config.Config
	GetLogger() logger.Logger
	GetRegistry() registry.Registry
	GetSelector() selector.Selector
	GetStore() store.Store
	GetTransport() transport.Transport
	GetEvents() events.Stream
}

type Config interface {
	GetService() micro.Service
	GetAuth() auth.Auth
	GetBroker() broker.Broker
	GetRegistry() registry.Registry
	GetLogger() logger.Logger
	GetWebService() web.Service
	SetService(micro.Service)
	SetRegistry(registry.Registry)
	SetAuth(auth.Auth)
	SetBroker(broker.Broker)
	SetClient(client.Client)
	SetLogger(logger.Logger)
	SetServer(server.Server)
	SetStore(store.Store)
	SetTransport(transport.Transport)
	SetWebService(web.Service)
	SetConfig(config.Config)
	SetSelector(selector.Selector)
	SetSync(syncx.Sync)
	SetEvents(events.Stream)
	Validate() error
}
