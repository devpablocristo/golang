package sdkgomicro

import (
	"fmt"
	"sync"

	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4/client"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

var (
	instance  ports.GrpcClient
	once      sync.Once
	initError error
)

type grpcImpl struct {
	client client.Client
}

func newGrpcClient(config ports.ConfigGrpcClient) (ports.GrpcClient, error) {
	once.Do(func() {
		clt, err := setupGrpcClient(config)
		if err != nil {
			initError = fmt.Errorf("error setting up gRPC client: %w", err)
			return
		}
		instance = &grpcImpl{
			client: clt,
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupGrpcClient(config ports.ConfigGrpcClient) (client.Client, error) {
	grpcClt := grpc.NewClient()

	return grpcClt, nil
}

func (c *grpcImpl) Client() client.Client {
	return c.client
}
