package core

import "api/internal/core/item"

// ItemUsecasePort define la interfaz para el caso de uso de elementos
type ItemUsecasePort interface {
	SaveItem(item.Item) error
	ListItems() (item.MapRepo, error)
	UpdateItem(item.Item) error
	DeleteItem(int) error
}
