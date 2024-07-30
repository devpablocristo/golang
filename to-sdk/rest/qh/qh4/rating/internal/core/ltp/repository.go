package ltp

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

func (r *Repository) GetLTP(ctx context.Context, ID string) error {
	return nil
}
