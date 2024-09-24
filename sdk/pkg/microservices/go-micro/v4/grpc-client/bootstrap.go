package sdkgomicro

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client/ports"
)

func Bootstrap() (ports.Client, error) {
	config := newConfig()

	return newClient(config)
}
