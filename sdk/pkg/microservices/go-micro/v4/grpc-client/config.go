package sdkgomicro

import "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"

type configGrpcClient struct{}

func newConfigGrpcClient() ports.ConfigGrpcClient {
	return &configGrpcClient{}
}
