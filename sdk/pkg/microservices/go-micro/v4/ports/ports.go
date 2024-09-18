package ports

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

type Service interface {
	StartRestServer() error
	StartGrpcService() error
	GetGrpcService() micro.Service
	GetRestServer() web.Service
}

type Config interface {
	GetConsulAddress() string
	Validate() error
}
