package domain

import (
	"errors"
	"time"
)

type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemRepositoryPort interface {
	SaveItem(Item) error
	ListItems() (MapRepo, error)
}

type MapRepo map[int]Item

var ErrNotFound = errors.New("not found")
