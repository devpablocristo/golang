package cin7

import (
	nimblecin7 "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7"
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type ShipmentReq struct {
	ShipmentID  string               `json:"shipment_id"`
	OrderID     string               `json:"order_id"`
	ShippedDate string               `json:"shipped_date"`
	Items       []nimblecin7.ItemReq `json:"items"`
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
