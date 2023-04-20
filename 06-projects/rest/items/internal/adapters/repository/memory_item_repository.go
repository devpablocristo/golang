package repository

import (
	"errors"
	"time"

	"github.com/mercadolibre/items/internal/entity"
)

var newID uint = 0

type itemRepository struct {
	db []entity.Item
}

func NewItemRepository() entity.ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetItems() ([]entity.Item, error) {
	return r.db, nil
}

func (r *itemRepository) GetItemByID(id uint) (entity.Item, error) {
	for _, item := range r.db {
		if item.ID == id {
			return item, nil
		}
	}

	return entity.Item{}, errors.New("item not found")
}

func (r *itemRepository) CheckItemByCode(code string) (bool, error) {
	for _, item := range r.db {
		if item.Code == code {
			return true, nil
		}
	}

	return false, nil
}

func (r *itemRepository) AddItem(item *entity.Item) error {
	createdAt := time.Now()
	newID = newID + 1

	item.ID = newID
	item.CreatedAt = createdAt
	item.UpdatedAt = createdAt
	r.db = append(r.db, *item)

	return nil
}
