package entity

import (
	"errors"
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
	SaveItem(Item Item) error
	GetItems() (MapRepo, error)
}

// este es implemente un alias para el mapa que se utilza en el repositorio inmemory

// este error es utilizado por mas de paquete por eso esta en este paquete pero es esto
// en realidad no es una buena implementacion por lo que debera mover
var ErrNotFound = errors.New("not found")
