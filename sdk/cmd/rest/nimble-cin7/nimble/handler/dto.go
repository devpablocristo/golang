package handler

import (
	nimblecin7 "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7"
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
)

type OrderReq struct {
	OrderID      string               `json:"order_id"`
	CustomerName string               `json:"customer_name"`
	Items        []nimblecin7.ItemReq `json:"items"`
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
