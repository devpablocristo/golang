package nimble

import (
	cin7 "github.com/devpablocristo/qh/events/internal/core/integration-nimble-cin7/cin7"
)

type Order struct {
	OrderID      string
	CustomerName string
	Items        []Item
}

type Item struct {
	ItemID   string
	Quantity int
}

// ConvertNimbleItemToCin7Item convierte un item de Nimble a un item de Cin7
func ToCin7Item(nimbleItem Item) cin7.Item {
	return cin7.Item{
		ItemID:   nimbleItem.ItemID,
		Quantity: nimbleItem.Quantity,
	}
}

// ConvertNimbleOrderToCin7Shipment convierte una orden de Nimble a un envío de Cin7
func toCin7Shipment(order Order) cin7.Shipment {
	// Convierte cada item de la orden de Nimble a un item de envío de Cin7
	items := make([]cin7.Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = ToCin7Item(item)
	}

	return cin7.Shipment{
		OrderID:     order.OrderID,
		ShippedDate: "", // Este campo se puede establecer según sea necesario
		Items:       items,
	}
}
