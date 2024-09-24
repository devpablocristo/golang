package authconn

import (
	"fmt"

	sdk "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8/ports"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type redisService struct {
	redis sdkports.Cache
}

func NewRedisService() (ports.RedisService, error) {
	c, err := sdk.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	return &redisService{
		redis: c,
	}, nil
}

func (c *redisService) Algo() error {
	return nil
}
