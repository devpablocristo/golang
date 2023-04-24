package repository

import (
	"fmt"

	entity "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/entity"
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

func (r *Repository) SaveItem(item entity.Item) error {
	if item.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[item.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", item.ID)
	}
	r.items[item.ID] = item
	return nil
}

func (r *Repository) GetAllItems() (entity.MapRepo, error) {
	return r.items, nil
}
