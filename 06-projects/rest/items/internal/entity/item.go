package entity

import (
	"fmt"
	"time"
)

type Item struct {
	ID          uint
	Code        string
	Description string
	Title       string
	Price       int
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemRepository interface {
	GetItems() ([]Item, error)
	GetItemByID(id uint) (Item, error)
	CheckItemByCode(code string) (bool, error)
	AddItem(item *Item) error
}

type ItemNotFound struct {
	Message string
}

func (e ItemNotFound) Error() string {
	return fmt.Sprintf("error: '%s'", e.Message)
}

type ItemAlreadyExist struct {
	Message string
}

func (e ItemAlreadyExist) Error() string {
	return fmt.Sprintf("error: '%s'", e.Message)
}
