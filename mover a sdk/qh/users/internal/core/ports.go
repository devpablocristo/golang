package core

import (
	"context"

	usr "github.com/devpablocristo/qh-users/internal/core/user"
)

type UseCasePort interface {
	GetUser(context.Context, string) (*usr.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*usr.InMemDB, error)
	UpdateUser(context.Context, *usr.User, string) error
	CreateUser(context.Context, *usr.User) error
	PublishMessage(context.Context, string) (string, error)
}
