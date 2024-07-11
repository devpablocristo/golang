package mongodb

import (
	"github.com/luka385/tinder-pets/domain"
)

type UserRepositoryPort interface {
	Create(*domain.User) error
	Update(*domain.User) error
	Delete(string) error
	GetByID(string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
}
