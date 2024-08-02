package nimble

import (
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type CachePort interface {
	CreateShipment(Order) (cin7.Shipment, error)
}
