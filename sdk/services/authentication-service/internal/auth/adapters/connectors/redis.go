package authconn

import (
	"log"

	sdk "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8/ports"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type redisService struct {
	redis sdkports.Cache
}

func NewRedisService() ports.RedisService {
	c, err := sdk.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return &redisService{
		redis: c,
	}
}

func (c *redisService) Algo() error {
	return nil
}
