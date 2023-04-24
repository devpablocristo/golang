package entity

import (
	"time"
)

type ID uint
type MapRepo map[ID]Item

// entidad Item
type Item struct {
	Code        string
	Title       string
	Description string
	Price       float64
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// esta es la interface qie utizaran cualquier repositorio que se implemente
type ItemRepository interface {
	SaveItem(Item Item) (Item, error)
	GetItems() (MapRepo, error)
}
