package redissetup

import (
	"github.com/spf13/viper"

	redisv8 "github.com/devpablocristo/golang/sdk/pkg/redis/v8"
)

func NewRedisInstance() (redisv8.RedisClientPort, error) {
	config := redisv8.RedisClientConfig{
		Address:  viper.GetString("REDIS_ADDRESS"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB"),
	}

	if err := redisv8.InitializeRedisClient(config); err != nil {
		return nil, err
	}

	return redisv8.GetRedisInstance()
}
