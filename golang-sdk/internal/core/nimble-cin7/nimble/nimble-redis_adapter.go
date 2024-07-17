package nimble

import (
	"time"

	cin7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/cin7"
)

type Redis struct{}

func NewNimbleRepository() RedisPort {
	return &Redis{}
}

func (r *Redis) CreateShipment(order Order) (cin7.Shipment, error) {
	shmnt := toCin7Shipment(order)
	shmnt.ShippedDate = time.Now().Format("2006-01-02")
	return shmnt, nil
}
