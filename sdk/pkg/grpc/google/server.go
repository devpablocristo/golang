package grpcpkg

import (
	"fmt"
	"net"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	serverInstance portspkg.GrpcServer
	onceServer     sync.Once
	errInitServer  error
)

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

// InitializeGrpcServer inicializa un servidor gRPC con la configuración proporcionada.
func InitializeGrpcServer(config portspkg.GrpcConfig) error {
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
		serverInstance = &grpcServer{server: server, listener: listener}
	})
	return errInitServer
}

// GetGrpcServerInstance devuelve la instancia del servidor gRPC.
func GetGrpcServerInstance() (portspkg.GrpcServer, error) {
	if serverInstance == nil {
		return nil, fmt.Errorf("grpc server is not initialized")
	}
	return serverInstance, nil
}

// Start inicia el servidor gRPC.
func (s *grpcServer) Start() error {
	return s.server.Serve(s.listener)
}

// Stop detiene el servidor gRPC.
func (s *grpcServer) Stop() error {
	s.server.GracefulStop()
	return s.listener.Close()
}

// RegisterService registra un servicio gRPC en el servidor.
func (s *grpcServer) RegisterService(serviceDesc *grpc.ServiceDesc, impl interface{}) {
	s.server.RegisterService(serviceDesc, impl)
}
