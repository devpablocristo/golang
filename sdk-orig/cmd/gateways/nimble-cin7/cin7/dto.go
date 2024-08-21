package cin7

import (
	"github.com/devpablocristo/golang/sdk/cmd/gateways/nimble-cin7/utils"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
	utils2 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/utils"
)

type ShipmentReq struct {
	ShipmentID  string          `json:"shipment_id"`
	OrderID     string          `json:"order_id"`
	ShippedDate string          `json:"shipped_date"`
	Items       []utils.ItemReq `json:"items"`
}

// ToCin7Shipment convierte un ShipmentReq en un cin7.Shipment
func ToCin7Shipment(shipmentReq ShipmentReq) cin7.Shipment {
	items := make([]utils2.Item, len(shipmentReq.Items)) // Usa utils.Item
	for i, itemReq := range shipmentReq.Items {
		items[i] = utils2.Item{ // Usa el tipo utils.Item
			ItemID:   itemReq.ItemID,
			Quantity: itemReq.Quantity,
		}
	}

	return cin7.Shipment{
		ShipmentID:  shipmentReq.ShipmentID,
		OrderID:     shipmentReq.OrderID,
		ShippedDate: shipmentReq.ShippedDate,
		Items:       items, // Usa el slice de utils.Item
	}
}
