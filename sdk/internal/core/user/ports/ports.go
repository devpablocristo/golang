package ports

import (
	"context"

	user "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

type Repository interface {
	SaveUser(context.Context, *user.User) error
	GetUser(context.Context, string) (*user.User, error)
	GetUserByUsername(context.Context, string) (*user.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*user.InMemDB, error)
	UpdateUser(context.Context, *user.User, string) error
}
