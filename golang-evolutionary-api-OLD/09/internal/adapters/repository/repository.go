package repository

import (
	"errors"
	"time"

	domain "items/internal/domain"
)

type Repository struct {
	items domain.MapRepo
}

func NewRepository() domain.ItemRepositoryPort {
	return &Repository{
		items: make(domain.MapRepo),
	}
}

func (r *Repository) SaveItem(item *domain.Item) (*domain.Item, error) {
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = item.CreatedAt
	id := domain.ID(len(r.items) + 1)
	r.items[id] = item

	return r.items[id], nil
}

func (r *Repository) GetItemByID(id domain.ID) (*domain.Item, error) {
	item, ok := r.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (r *Repository) GetItemByCode(code string) (*domain.Item, error) {
	for _, item := range r.items {
		if item.Code == code {
			return item, errors.New("existing code")
		}
	}

	return nil, nil
}

func (r *Repository) GetAllItems() (domain.MapRepo, error) {
	return r.items, nil
}
