package item

import (
	"fmt"
)

type Repository struct {
	items MapRepo
}

func NewRepository() ItemRepositoryPort {
	return &Repository{
		items: make(MapRepo),
	}
}

func (r *Repository) SaveItem(it Item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", it.ID)
	}
	r.items[it.ID] = it
	return nil
}

func (r *Repository) ListItems() (MapRepo, error) {
	return r.items, nil
}
