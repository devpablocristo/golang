package handler

import (
	nimble "github.com/devpablocristo/qh/events/internal/core/nimble"
)

type order struct {
	OrderID      string `json:"order_id"`
	CustomerName string `json:"customer_name"`
	Items        []item `json:"items"`
}

type item struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type Cin7Shipment struct {
	ShipmentID  string `json:"shipment_id"`
	OrderID     string `json:"order_id"`
	ShippedDate string `json:"shipped_date"`
	Items       []item `json:"items"`
}

func itemsToDomain(items []item) []nimble.Item {
	domainItems := make([]nimble.Item, len(items))
	for i, item := range items {
		domainItems[i] = nimble.Item{
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
	}
	return domainItems
}
