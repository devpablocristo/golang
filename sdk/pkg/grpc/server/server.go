package sdkcserver

import (
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	ports "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

var (
	serverInstance ports.Server
	serverOnce     sync.Once
	serverInitErr  error
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func newServer(config ports.Config) (ports.Server, error) {
	serverOnce.Do(func() {
		var opts []grpc.ServerOption
		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				serverInitErr = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.Creds(creds))
		}

		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()))
		if err != nil {
			serverInitErr = fmt.Errorf("failed to listen: %v", err)
			return
		}

		server := grpc.NewServer(opts...)
		reflection.Register(server) // Registro de reflexión gRPC

		serverInstance = &Server{server: server, listener: listener}
	})
	return serverInstance, serverInitErr
}

// GetServerInstance devuelve la instancia de servidor gRPC
func GetServerInstance() (ports.Server, error) {
	if serverInstance == nil {
		return nil, fmt.Errorf("gRPC server is not initialized")
	}
	return serverInstance, nil
}

func (s *Server) Start() error {
	return s.server.Serve(s.listener)
}

func (s *Server) Stop() error {
	s.server.GracefulStop()
	return s.listener.Close()
}

func (s *Server) RegisterService(serviceDesc *grpc.ServiceDesc, impl any) {
	s.server.RegisterService(serviceDesc, impl)
}
