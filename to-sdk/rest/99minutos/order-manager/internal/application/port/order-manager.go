package port

import (
	"context"

	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

//go:generate mockgen -source=./order-manager.go -destination=../../mocks/order-manager_mock.go -package=mocks
type OrderManager interface {
	CreateOrder(context.Context, *domain.Order) (*domain.Order, error)
}
