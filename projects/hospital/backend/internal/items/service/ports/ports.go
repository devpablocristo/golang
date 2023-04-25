package ports

import (
	"context"

	"github.com/devpablocristo/golang/hex-arch/backend/internal/items/domain"
)

//go:generate mockgen -source=./services.go -destination=../test/mocks/item_service_mock.go -package=mocks
type ItemService interface {
	CreateItem(ctx context.Context, item domain.Item) (*domain.Item, error)
	GetItemByID(ctx context.Context, itemID uint) (*domain.Item, error)
}
