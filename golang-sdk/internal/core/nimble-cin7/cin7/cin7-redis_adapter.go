package cin7

import (
	"context"

	redisv8 "github.com/devpablocristo/qh/events/pkg/redis/v8"
)

type RedisRepository struct {
	redisInst redisv8.RedisClientPort
}

func NewRedisRepository(inst redisv8.RedisClientPort) RedisPort {
	return &RedisRepository{
		redisInst: inst,
	}
}

func (r *RedisRepository) SaveShipment(shipment Shipment) error {
	ctx := context.Background()
	return r.redisInst.Client().Set(ctx, shipment.ShipmentID, shipment, 0).Err()
}
