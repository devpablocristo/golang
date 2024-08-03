package nimble

import (
	dtoshared "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/shared"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/shared"
)

// OrderReq representa la estructura del request para una orden de Nimble
type OrderReq struct {
	OrderID      string              `json:"order_id"`
	CustomerName string              `json:"customer_name"`
	Items        []dtoshared.ItemReq `json:"items"`
}

// ToNimbleOrder convierte un OrderReq en un nimble.Order
func ToNimbleOrder(orderReq OrderReq) nimble.Order {
	items := make([]shared.Item, len(orderReq.Items)) // Usa core.Item para asegurar la consistencia
	for i, itemReq := range orderReq.Items {
		items[i] = shared.Item{
			ItemID:   itemReq.ItemID,
			Quantity: itemReq.Quantity,
		}
	}

	return nimble.Order{
		OrderID:      orderReq.OrderID,
		CustomerName: orderReq.CustomerName,
		Items:        items, // Asigna directamente los items de tipo core.Item
	}
}
