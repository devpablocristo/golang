package ports

import "go-micro.dev/v4/server"

type Server interface {
	GetServer() server.Server
}

type Config interface {
	GetGrpcServerName() string
	GetGrpcServerHost() string
	GetGrpcServerPort() int
	Validate() error
}
