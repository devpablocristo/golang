package cin7

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) RedisPort {
	return &Redis{client: client}
}

func (r *Redis) SaveShipment(shipment Shipment) error {
	ctx := context.Background()
	return r.client.Set(ctx, shipment.ShipmentID, shipment, 0).Err()
}
