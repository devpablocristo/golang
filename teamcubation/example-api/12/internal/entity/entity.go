package entity

import (
	"time"
)

type ID uint
type MapRepo map[ID]*Item

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

// esta es la interface del repositorio, o sea, el conector del repositorio con el resto de la app
type ItemRepositoryPort interface {
	SaveItem(*Item) (*Item, error)
	GetAllItems() (MapRepo, error)
	CheckItemByCode(string) (bool, error)
	GetItemByID(ID) (*Item, error)
}
