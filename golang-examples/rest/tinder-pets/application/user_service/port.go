package application

import (
	"github.com/luka385/tinder-pets/domain"
)

type UserServicePort interface {
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	GetUserByID(id string) (*domain.User, error)
	GetAllUser() ([]*domain.User, error)
}
