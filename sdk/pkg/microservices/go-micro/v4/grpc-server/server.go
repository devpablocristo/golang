package sdkgomicro

import (
	"fmt"
	"sync"

	"go-micro.dev/v4/server"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server/ports"
)

var (
	instance  ports.Server
	once      sync.Once
	initError error
)

type Server struct {
	server server.Server
}

func newServer(config ports.Config) (ports.Server, error) {
	once.Do(func() {
		srv, err := setupServer(config)
		if err != nil {
			initError = fmt.Errorf("error setting up  server: %w", err)
			return
		}
		instance = &Server{
			server: srv,
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupServer(config ports.Config) (server.Server, error) {
	s := server.NewServer(
		server.Name(config.GetServerName()),
		server.Id(config.GetServerID()),
		server.Address(fmt.Sprintf("%s:%d", config.GetServerHost(), config.GetServerPort())),
	)

	return s, nil
}

func (s *Server) GetServer() server.Server {
	return s.server
}
