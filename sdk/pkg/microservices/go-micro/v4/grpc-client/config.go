package sdkgomicro

import "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client/ports"

type config struct{}

func newConfig() ports.Config {
	return &config{}
}
