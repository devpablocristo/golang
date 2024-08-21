package usercoreports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

type Repository interface {
	SaveUser(context.Context, *entities.User) error
	GetUser(context.Context, string) (*entities.User, error)
	GetUserUUID(context.Context, string, string) (string, error)
	GetUserByUsername(context.Context, string) (*entities.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*entities.InMemDB, error)
	UpdateUser(context.Context, *entities.User, string) error
}

type UserUseCases interface {
	GetUser(context.Context, string) (*entities.User, error)
	GetUserUUID(context.Context, string, string) (string, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*entities.InMemDB, error)
	UpdateUser(context.Context, *entities.User, string) error
	CreateUser(context.Context, *entities.User) error
	PublishMessage(context.Context, string) (string, error)
}
