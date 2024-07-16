package cin7

// import (
// 	nimble "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7/nimble"
// )

type Shipment struct {
	ShipmentID  string
	OrderID     string
	ShippedDate string
	Items       []Item
}

type Item struct {
	ItemID   string
	Quantity int
}
