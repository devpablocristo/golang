package repository

import (
	"errors"
	"time"

	entity "items/internal/entity"
)

// el campo items es del type que maneja el repositorio
type repository struct {
	items entity.MapRepo
}

// de nuevo, aqui el tipo retornado utiliza una interface
func NewRepository() entity.ItemRepository {
	return &repository{
		items: make(entity.MapRepo),
	}
}

func (r *repository) SaveItem(item *entity.Item) (*entity.Item, error) {
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = item.CreatedAt
	id := entity.ID(len(r.items) + 1)
	r.items[id] = item

	return r.items[id], nil
}

func (r *repository) GetItemByID(id entity.ID) (*entity.Item, error) {
	item, ok := r.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (r *repository) CheckItemByCode(code string) (bool, error) {
	for _, item := range r.items {
		if item.Code == code {
			return true, errors.New("existing code")
		}
	}

	return false, nil
}

func (r *repository) GetAllItems() (entity.MapRepo, error) {
	return r.items, nil
}
