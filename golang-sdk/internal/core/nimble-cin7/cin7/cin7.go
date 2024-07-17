package cin7

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
