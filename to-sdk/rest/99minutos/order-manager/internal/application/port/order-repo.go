package port

import (
	"context"

	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

//go:generate mockgen -source=./order-repo.go -destination=../../mocks/order-repo_mock.go -package=mocks
type OrderRepo interface {
	Create(context.Context, *domain.Order) error
	Read(context.Context, string) (*domain.Order, error)
	List(context.Context) map[string]*domain.Order
	Delete(context.Context, string) error
}
