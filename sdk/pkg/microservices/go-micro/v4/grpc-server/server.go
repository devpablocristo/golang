package sdkgomicro

import (
	"fmt"
	"sync"

	"go-micro.dev/v4/server"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

var (
	instance  ports.GrpcServer
	once      sync.Once
	initError error
)

type grpcServer struct {
	server server.Server
}

func newGrpcServer(config ports.ConfigGrpcServer) (ports.GrpcServer, error) {
	once.Do(func() {
		srv, err := setupGrpcServer(config)
		if err != nil {
			initError = fmt.Errorf("error setting up gRPC server: %w", err)
			return
		}
		instance = &grpcServer{
			server: srv,
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcServer(config ports.ConfigGrpcServer) (server.Server, error) {
	grpcSrv := server.NewServer(
		server.Address(fmt.Sprintf("%s:%d", config.GetGrpcServerHost(), config.GetGrpcServerPort())),
	)

	return grpcSrv, nil
}

func (s *grpcServer) Server() server.Server {
	return s.server
}
