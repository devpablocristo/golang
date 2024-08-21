package redisv8

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	instance RedisClientPort
	once     sync.Once
	errInit  error
)

type RedisClientPort interface {
	Client() *redis.Client
	Close()
}

type RedisClient struct {
	client *redis.Client
}

func InitializeRedisClient(config RedisClientConfig) error {
	once.Do(func() {
		client := &RedisClient{}
		errInit = client.connect(config)
		if errInit != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return errInit
}

func GetRedisInstance() (RedisClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("redis client is not initialized")
	}
	return instance, nil
}

func (client *RedisClient) connect(config RedisClientConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return err
	}
	client.client = rdb
	return nil
}

func (client *RedisClient) Close() {
	if client.client != nil {
		client.client.Close()
	}
}

func (client *RedisClient) Client() *redis.Client {
	return client.client
}