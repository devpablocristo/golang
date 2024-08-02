package nimble

import (
	nc "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/shared"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
)

// OrderReq representa la estructura del request para una orden de Nimble
type OrderReq struct {
	OrderID      string       `json:"order_id"`
	CustomerName string       `json:"customer_name"`
	Items        []nc.ItemReq `json:"items"`
}

// ToNimbleOrder convierte un OrderReq en un nimble.Order
func ToNimbleOrder(orderReq OrderReq) nimble.Order {
	items := make([]core.Item, len(orderReq.Items)) // Usa core.Item para asegurar la consistencia
	for i, itemReq := range orderReq.Items {
		items[i] = core.Item{
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
