package nimble

import (
	"github.com/devpablocristo/golang/sdk/cmd/gateways/nimble-cin7/utils"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
	utils2 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/utils"
)
type ItemReq struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

// OrderReq representa la estructura del request para una orden de Nimble
type OrderReq struct {
	OrderID      string           `json:"order_id"`
	CustomerName string           `json:"customer_name"`
	Items        []utils.ItemReq `json:"items"`
}

// ToNimbleOrder convierte un OrderReq en un nimble.Order
func ToNimbleOrder(orderReq OrderReq) nimble.Order {
	items := make([]utils2.Item, len(orderReq.Items)) // Usa core.Item para asegurar la consistencia
	for i, itemReq := range orderReq.Items {
		items[i] = utils2.Item{
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
