package nimble

import (
	"time"

	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
	redisv8 "github.com/devpablocristo/golang/sdk/pkg/redis/v8"
)

type redisRepository struct {
	redisInst redisv8.RedisClientPort
}

func NewRedisRepository(inst redisv8.RedisClientPort) CacheRepository {
	return &redisRepository{
		redisInst: inst,
	}
}

func (r *redisRepository) CreateShipment(order Order) (cin7.Shipment, error) {
	shmnt := toCin7Shipment(order)
	shmnt.ShippedDate = time.Now().Format("2006-01-02")
	return shmnt, nil
}

// toCin7Shipment convierte una orden de Nimble a un envío de Cin7
func toCin7Shipment(order Order) cin7.Shipment {
	return cin7.Shipment{
		OrderID:     order.OrderID,
		ShippedDate: "",          // Este campo se puede establecer según sea necesario
		Items:       order.Items, // Usa directamente los items de shared
	}
}
