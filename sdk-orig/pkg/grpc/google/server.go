package ggrpcgpkg

import (
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
)

var (
	serverInstance portspkg.GgrpcServer
	onceServer     sync.Once
	errInitServer  error
)

type ggrpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func InitializeGgrpcServer(config portspkg.GgrpcConfig) error {
	onceServer.Do(func() {
		var opts []grpc.ServerOption
		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				errInitServer = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.Creds(creds))
		}

		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()))
		if err != nil {
			errInitServer = fmt.Errorf("failed to listen: %v", err)
			return
		}

		server := grpc.NewServer(opts...)
		reflection.Register(server) // Registro de reflexión gRPC

		serverInstance = &ggrpcServer{server: server, listener: listener}
	})
	return errInitServer
}

// GetGrpcServerInstance devuelve la instancia del servidor gRPC.
func GetGgrpcServerInstance() (portspkg.GgrpcServer, error) {
	if serverInstance == nil {
		return nil, fmt.Errorf("grpc server is not initialized")
	}
	return serverInstance, nil
}

// Start inicia el servidor gRPC.
func (s *ggrpcServer) Start() error {
	return s.server.Serve(s.listener)
}

// Stop detiene el servidor gRPC.
func (s *ggrpcServer) Stop() error {
	s.server.GracefulStop()
	return s.listener.Close()
}

// RegisterService registra un servicio gRPC en el servidor.
func (s *ggrpcServer) RegisterService(serviceDesc *grpc.ServiceDesc, impl any) {
	s.server.RegisterService(serviceDesc, impl)
}
