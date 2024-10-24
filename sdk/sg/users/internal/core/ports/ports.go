package ports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UseCases interface {
	CheckUserStatus(context.Context, string) (bool, error)
	CreateUser(ctx context.Context, user *entities.User) error
}

type Repository interface {
	CreateUser(context.Context, *entities.User) error
	FindUserByUUID(context.Context, string) (*entities.User, error)
	FindUserByCuit(context.Context, string) (*entities.User, error)
	UpdateUser(context.Context, *entities.User) error
	SoftDeleteUser(context.Context, string) error
}
