package repository

import (
	entity "Items/internal/domain"
)

type repository struct {
}

func NewRepository() entity.Repository {
	return repository{}
}

func (r repository) AddItem(item entity.Item) (entity.Item, error) {

	return entity.Item{}, nil
}
