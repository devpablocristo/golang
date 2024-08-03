package report

import (
	"context"
)

type mapRepository struct {
	db *InMemDB
}

func NewMapRepository() RepositoryPort {
	db := make(InMemDB)
	return &mapRepository{
		db: &db,
	}
}

func (r *mapRepository) CreateReport(ctx context.Context) error {
	return nil
}
