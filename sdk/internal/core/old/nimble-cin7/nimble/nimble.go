package nimble

import (
	"github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/utils"
)

type Order struct {
	OrderID      string
	CustomerName string
	Items        []utils.Item
}
