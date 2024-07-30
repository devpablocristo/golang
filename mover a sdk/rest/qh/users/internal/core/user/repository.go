package usr

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Repository struct {
	db *InMemDB
}

func NewRepository() RepositoryPort {
	db := make(InMemDB) // Inicializa el mapa.
	return &Repository{
		db: &db,
	}
}

func (r *Repository) GetUser(ctx context.Context, ID string) (*User, error) {
	usr, exists := (*r.db)[ID]
	if !exists {
		return nil, errors.New("usr not found")
	}
	return usr, nil
}

func (r *Repository) DeleteUser(ctx context.Context, ID string) error {
	if _, exists := (*r.db)[ID]; !exists {
		return errors.New("usr not found")
	}
	delete(*r.db, ID)
	return nil
}

func (r *Repository) ListUsers(ctx context.Context) (*InMemDB, error) {
	return r.db, nil
}

func (r *Repository) UpdateUser(ctx context.Context, usr *User, ID string) error {
	if _, exists := (*r.db)[ID]; !exists {
		return errors.New("usr not found")
	}
	(*r.db)[ID] = usr
	return nil
}

func (r *Repository) CreateUser(ctx context.Context, usr *User) error {
	ID := uuid.New().String()
	(*r.db)[ID] = usr
	return nil
}
