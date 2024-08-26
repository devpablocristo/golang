package pkggomicroports

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

type Service interface {
	StartWebService() error
	StartRcpService() error
	GetRcpService() micro.Service
	GetWebService() web.Service
}

type Config interface {
	GetRcpServiceName() string
	GetWebServiceName() string
	GetRcpServiceAddress() string
	GetWebServiceAddress() string
	GetConsulAddress() string
	Validate() error
}
