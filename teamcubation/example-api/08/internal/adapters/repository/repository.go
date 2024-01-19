package repository

import (
	"fmt"

	"items/internal/domain"
)

type Repository struct {
	items domain.MapRepo
}

func NewRepository() domain.ItemRepositoryPort {
	return &Repository{
		items: make(domain.MapRepo),
	}
}

func (r *Repository) SaveItem(item domain.Item) error {
	if item.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[item.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", item.ID)
	}
	r.items[item.ID] = item
	return nil
}

func (r *Repository) GetAllItems() (domain.MapRepo, error) {
	return r.items, nil
}
