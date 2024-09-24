package sdkgomicro

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func BootstrapGrpcClient() (ports.GrpcClient, error) {
	config := newConfigGrpcClient()

	return newGrpcClient(config)
}
