package inmemoryr

import (
	"errors"
	"time"

	domain "items/internal/domain"
)

type InMemory struct {
	items domain.MapRepo
}

func NewInMemory() domain.ItemRepositoryPort {
	return &InMemory{
		items: make(domain.MapRepo),
	}
}

func (r *InMemory) SaveItem(item *domain.Item) (*domain.Item, error) {
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = item.CreatedAt
	id := domain.ID(len(r.items) + 1)
	r.items[id] = item

	return r.items[id], nil
}

func (r *InMemory) GetItem(id domain.ID) (*domain.Item, error) {
	item, ok := r.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (r *InMemory) GetItemByCode(code string) (*domain.Item, error) {
	for _, item := range r.items {
		if item.Code == code {
			return item, errors.New("existing code")
		}
	}

	return nil, nil
}

func (r *InMemory) GetAllItems() (domain.MapRepo, error) {
	return r.items, nil
}
