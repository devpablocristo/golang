package handler

import (
	cin7 "github.com/devpablocristo/qh/events/internal/core/integration-nimble-cin7/cin7"
	nimble "github.com/devpablocristo/qh/events/internal/core/integration-nimble-cin7/nimble"
)

type OrderReq struct {
	OrderID      string    `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	Items        []ItemReq `json:"items"`
}

type ItemReq struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type ShipmentReq struct {
	ShipmentID  string    `json:"shipment_id"`
	OrderID     string    `json:"order_id"`
	ShippedDate string    `json:"shipped_date"`
	Items       []ItemReq `json:"items"`
}

func ToNimbleOrder(orderReq OrderReq) nimble.Order {
	items := make([]nimble.Item, len(orderReq.Items))
	for i, itemReq := range orderReq.Items {
		items[i] = nimble.Item{
			ItemID:   itemReq.ItemID,
			Quantity: itemReq.Quantity,
		}
	}

	return nimble.Order{
		OrderID:      orderReq.OrderID,
		CustomerName: orderReq.CustomerName,
		Items:        items,
	}
}

func ToCin7Shipment(shipmentReq ShipmentReq) cin7.Shipment {
	items := make([]cin7.Item, len(shipmentReq.Items))
	for i, itemReq := range shipmentReq.Items {
		items[i] = cin7.Item{
			ItemID:   itemReq.ItemID,
			Quantity: itemReq.Quantity,
		}
	}

	return cin7.Shipment{
		ShipmentID:  shipmentReq.ShipmentID,
		OrderID:     shipmentReq.OrderID,
		ShippedDate: shipmentReq.ShippedDate,
		Items:       items,
	}
}
