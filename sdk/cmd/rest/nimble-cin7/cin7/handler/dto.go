package cin7

import (
	dtoshared "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/shared"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/shared"
)

type ShipmentReq struct {
	ShipmentID  string              `json:"shipment_id"`
	OrderID     string              `json:"order_id"`
	ShippedDate string              `json:"shipped_date"`
	Items       []dtoshared.ItemReq `json:"items"`
}

// ToCin7Shipment convierte un ShipmentReq en un cin7.Shipment
func ToCin7Shipment(shipmentReq ShipmentReq) cin7.Shipment {
	items := make([]shared.Item, len(shipmentReq.Items)) // Usa shared.Item
	for i, itemReq := range shipmentReq.Items {
		items[i] = shared.Item{ // Usa el tipo shared.Item
			ItemID:   itemReq.ItemID,
			Quantity: itemReq.Quantity,
		}
	}

	return cin7.Shipment{
		ShipmentID:  shipmentReq.ShipmentID,
		OrderID:     shipmentReq.OrderID,
		ShippedDate: shipmentReq.ShippedDate,
		Items:       items, // Usa el slice de shared.Item
	}
}
