package cin7

import "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/utils"

type Shipment struct {
	ShipmentID  string
	OrderID     string
	ShippedDate string
	Items       []utils.Item // Items de tipo shared.Item
}
