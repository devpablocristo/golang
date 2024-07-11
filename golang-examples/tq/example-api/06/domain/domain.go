package domain

import (
	"errors"
	"time"
)

// entidad Item
type Item struct {
	ID          int
	Code        string
	Title       string
	Description string
	Price       float64
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemRepositoryPort interface {
	SaveItem(Item Item) error
	GetAllItems() (MapRepo, error)
}

type MapRepo map[int]Item

var ErrNotFound = errors.New("not found")
