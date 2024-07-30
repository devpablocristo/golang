package repository

import (
	"fmt"

	"api/internal/domain"
)

type Repository struct {
	items domain.MapRepo
}

func NewRepository() domain.ItemRepositoryPort {
	return &Repository{
		items: make(domain.MapRepo),
	}
}

func (r *Repository) SaveItem(it domain.Item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", it.ID)
	}
	r.items[it.ID] = it
	return nil
}

func (r *Repository) ListItems() (domain.MapRepo, error) {
	return r.items, nil
}
