package ports

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
)

type Service interface {
	SetWebRouter(router interface{}) error
}

type Config interface {
	GetGrpcServiceName() string
	GetWebServerName() string
	GetWebServerHost() string
	GetWebServerPort() int
	GetWebServerAddress() string
	GetConsulAddress() string
	GetGrpcServerHost() string
	GetGrpcServerPort() int
	GetWebRouter() interface{}
	Validate() error
}

type WebServer interface {
	SetWebRouter(router interface{}) error
}

type ConfigWebServer interface {
	GetWebServerName() string
	GetWebServerHost() string
	GetWebServerPort() int
	GetWebServerAddress() string
	GetConsulAddress() string
	GetWebRouter() interface{}
	Validate() error
}

type GrpcClient interface {
	Client() client.Client
}

type ConfigGrpcClient interface{}

type GrpcServer interface {
	Server() server.Server
}

type ConfigGrpcServer interface {
	GetGrpcServerName() string
	GetGrpcServerHost() string
	GetGrpcServerPort() int
	Validate() error
}

type GrpcService interface {
	Run() error
	Service() micro.Service
}

type ConfigGrpcService interface {
	GetServiceName() string
	GetServer() server.Server
	GetClient() client.Client
	GetConsulAddress() string
	Validate() error
}
