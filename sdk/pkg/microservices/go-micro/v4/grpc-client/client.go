package sdkgomicro

import (
	"sync"

	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4/client"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client/ports"
)

var (
	instance  ports.Client
	once      sync.Once
	initError error
)

type grpcClient struct {
	client client.Client
}

func newClient(config ports.Config) (ports.Client, error) {
	once.Do(func() {
		instance = &grpcClient{
			client: setupClient(config),
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupClient(config ports.Config) client.Client {
	grpcClt := grpc.NewClient()

	return grpcClt
}

func (c *grpcClient) GetClient() client.Client {
	return c.client
}
