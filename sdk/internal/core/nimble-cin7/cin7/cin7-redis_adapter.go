package cin7

import (
	"context"
	"encoding/json"

	redisv8 "github.com/devpablocristo/golang/sdk/pkg/redis/v8"
)

type RedisRepository struct {
	redisInst redisv8.RedisClientPort
}

func NewRedisRepository(inst redisv8.RedisClientPort) CachePort {
	return &RedisRepository{
		redisInst: inst,
	}
}

func (r *RedisRepository) SaveShipment(shipment Shipment) error {
	ctx := context.Background()
	data, err := json.Marshal(shipment)
	if err != nil {
		return err
	}
	return r.redisInst.Client().Set(ctx, shipment.ShipmentID, data, 0).Err()
}
