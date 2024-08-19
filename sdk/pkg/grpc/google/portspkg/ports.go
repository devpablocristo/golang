package portspkg

import (
	"context"

	"google.golang.org/grpc"
)

type GgrpcConfig interface {
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

type GgrpcClient interface {
	InvokeMethod(ctx context.Context, method string, request, response any) error
	Close() error
}

type GgrpcServer interface {
	Start() error
	Stop() error
	RegisterService(serviceDesc *grpc.ServiceDesc, impl any)
}
