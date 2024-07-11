package core

import "api/internal/core/item"

type ItemUsecasePort interface {
	SaveItem(item.Item) error
	ListItems() (item.MapRepo, error)
}
