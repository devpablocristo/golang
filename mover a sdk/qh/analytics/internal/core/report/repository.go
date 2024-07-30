package report

import (
	"context"
)

// This file is only an example of a starting point for a repository

type Repository struct {
	db *InMemDB
}

func NewRepository() RepositoryPort {
	db := make(InMemDB)
	return &Repository{
		db: &db,
	}
}

func (r *Repository) CreateReport(ctx context.Context) error {
	return nil
}
