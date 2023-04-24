package repository

import (
	"errors"
	"time"

	entity "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/entity"
)

// el campo items es del type que maneja el repositorio
type Repository struct {
	items entity.MapRepo
}

// de nuevo, aqui el tipo retornado utiliza una interface
func NewRepository() entity.ItemRepository {
	return &Repository{
		items: make(entity.MapRepo),
	}
}

func (r *Repository) SaveItem(item entity.Item) (entity.Item, error) {
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	id := entity.ID(len(r.items) + 1)
	r.items[id] = item

	return r.items[id], nil
}

func (r *Repository) GetItemByID(id entity.ID) (entity.Item, error) {
	item, ok := r.items[id]
	if !ok {
		return entity.Item{}, errors.New("item not found")
	}
	return item, nil
}

func (r *Repository) GetItems() (entity.MapRepo, error) {
	return r.items, nil
}
