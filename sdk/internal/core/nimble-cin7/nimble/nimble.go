package nimble

import (
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/shared"
)

type Order struct {
	OrderID      string
	CustomerName string
	Items        []shared.Item
}
