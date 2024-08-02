package cin7

import "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/shared"

type Shipment struct {
	ShipmentID  string
	OrderID     string
	ShippedDate string
	Items       []shared.Item // Items de tipo shared.Item
}
