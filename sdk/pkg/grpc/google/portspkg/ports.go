package portspkg

import (
	"context"

	"google.golang.org/grpc"
)

type GrpcConfig interface {
	GetHost() string
	SetHost(host string)
	GetPort() int
	SetPort(port int)
	GetTLSConfig() *TLSConfig
	SetTLSConfig(tlsConfig *TLSConfig)
	Validate() error
}

type TLSConfig struct {
	CertFile string
	KeyFile  string
	CAFile   string
}

type GrpcClient interface {
	InvokeMethod(ctx context.Context, method string, request, response interface{}) error
	Close() error
} 

type GrpcServer interface {
	Start() error
	Stop() error
	RegisterService(serviceDesc *grpc.ServiceDesc, impl interface{})
}
