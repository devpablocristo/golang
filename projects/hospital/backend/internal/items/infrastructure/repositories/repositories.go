package ports

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"

	item "github.com/devpablocristo/golang/hex-arch/backend/internal/items/domain"
)

//go:generate mockgen -source=./repositories.go -destination=../test/mocks/item_repository_mock.go -package=mocks
type ItemRepository interface {
	SaveItem(ctx context.Context, a *item.Item) error
	GetItemByID(ctx context.Context, id uint) (*item.Item, error)
}
