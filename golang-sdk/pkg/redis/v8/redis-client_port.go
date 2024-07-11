package redisv8

import "github.com/go-redis/redis/v8"

type RedisClientPort interface {
	Client() *redis.Client
	Close()
}
