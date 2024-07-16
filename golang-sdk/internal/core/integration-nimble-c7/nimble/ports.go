package nimble

import (
	cin7 "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7/c7"
)

type RedisPort interface {
	CreateShipment(Order) (cin7.Shipment, error)
}
