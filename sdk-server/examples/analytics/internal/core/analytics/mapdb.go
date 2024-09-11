package report

import (
	"context"

	"github.com/devpablocristo/golang/sdk/examples/analytics/internal/core/analytics/entities"
	"github.com/devpablocristo/golang/sdk/examples/analytics/internal/core/analytics/ports"
)

type mapRepository struct {
	db *entities.InMemDB
}

func NewMapRepository() ports.Repository {
	db := make(entities.InMemDB)
	return &mapRepository{
		db: &db,
	}
}

func (r *mapRepository) CreateReport(ctx context.Context) error { //-
	return nil
}
