package ports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

// Repository define las operaciones básicas que cualquier repositorio de usuarios debe implementar
type Repository interface {
	SaveUser(context.Context, *entities.User) error
	GetUser(context.Context, string) (*entities.User, error)
	GetUserByUsername(context.Context, string) (*entities.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*entities.InMemDB, error)
	UpdateUser(context.Context, *entities.User, string) error
}
