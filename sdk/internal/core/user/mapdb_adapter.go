package user

import (
	"context"

	"github.com/google/uuid"
)

type MapDbRepository struct {
	db *InMemDB
}

func NewMapDbRepository() RepositoryPort {
	db := make(InMemDB)
	return &MapDbRepository{
		db: &db,
	}
}

func (r *MapDbRepository) SaveUser(ctx context.Context, usr *User) error {
	ID := uuid.New().String()
	(*r.db)[ID] = usr
	return nil
}

func (r *MapDbRepository) GetUser(context.Context, string) (*User, error) {
	return nil, nil

}

func (r *MapDbRepository) GetUserByUsername(context.Context, string) (*User, error) {
	return nil, nil
}

// func (r *MapDbRepository) DeleteUser(ctx context.Context, ID string) error {
// 	if _, exists := (*r.db)[ID]; !exists {
// 		return errors.New("usr not found")
// 	}
// 	delete(*r.db, ID)
// 	return nil
// }

// func (r *MapDbRepository) ListUsers(ctx context.Context) (*InMemDB, error) {
// 	return r.db, nil
// }

// func (r *MapDbRepository) UpdateUser(ctx context.Context, usr *User, ID string) error {
// 	if _, exists := (*r.db)[ID]; !exists {
// 		return errors.New("usr not found")
// 	}
// 	(*r.db)[ID] = usr
// 	return nil
// }
