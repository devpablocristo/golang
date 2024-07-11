package redisv8

import "fmt"

type RedisClientConfig struct {
	Address  string
	Password string
	DB       int
}

func (config RedisClientConfig) Validate() error {
	if config.Address == "" {
		return fmt.Errorf("REDIS_ADDRESS is required")
	}
	if config.DB < 0 {
		return fmt.Errorf("REDIS_DB must be a non-negative integer")
	}
	return nil
}
