package portspkg

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
	syncm "go-micro.dev/v4/sync"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/web"
)

type GoMicroService interface {
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

type GoMicroConfig interface {
	GetName() string
	GetVersion() string
	GetAddress() string
	SetRegistry(reg registry.Registry)
	SetAuth(auth auth.Auth)
	SetBroker(broker broker.Broker)
	SetClient(client client.Client)
	SetLogger(logger logger.Logger)
	SetServer(server server.Server)
	SetStore(store store.Store)
	SetTransport(transport transport.Transport)
	SetWebService(webService web.Service)
	SetConfig(conf config.Config)
	SetSelector(selector selector.Selector)
	SetSync(sync syncm.Sync)
	SetEvents(events events.Stream)
	Validate() error
}
