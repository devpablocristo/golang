package core

import (
	cin7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/cin7"
)

type Cin7UseCasePort interface {
	UpdateShipment(cin7.Shipment) error
}