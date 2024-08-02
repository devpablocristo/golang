package nimble

import (
	"time"

	redisv8 "github.com/devpablocristo/golang/sdk/pkg/redis/v8"

	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type RedisRepository struct {
	redisInst redisv8.RedisClientPort
}

func NewRedisRepository(inst redisv8.RedisClientPort) RedisPort {
	return &RedisRepository{
		redisInst: inst,
	}
}

func (r *RedisRepository) CreateShipment(order Order) (cin7.Shipment, error) {
	shmnt := toCin7Shipment(order)
	shmnt.ShippedDate = time.Now().Format("2006-01-02")
	return shmnt, nil
}
