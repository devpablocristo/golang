package core

import (
	cin7 "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7/c7"
)

type Cin7UseCasePort interface {
	UpdateShipment(cin7.Shipment) error
}
