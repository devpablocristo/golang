package gomicroports

import (
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/web"
)

type GoMicroClient interface {
	Start() error
	Stop() error
	GetService() interface{}
	GetWebService() web.Service
	GetGrpcClient() client.Client
	GetGrpcServer() server.Server
}

type GoMicroConfig interface {
	GetName() string
	SetName(string)
	GetVersion() string
	SetVersion(string)
	GetAddress() string
	SetAddress(string)
	GetRegistry() interface{}
	SetRegistry(interface{})
	Validate() error
}
