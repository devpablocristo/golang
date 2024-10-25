package ports

import (
	"context"

	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UseCases interface {
	// CheckUserStatus(context.Context, string) (bool, error)
	CreateUser(context.Context, *dto.UserDto) (string, error)
}

type Repository interface {
	CreateUser(context.Context, *entities.User) error
	FindUserByUUID(context.Context, string) (*entities.User, error)
	FindUserByCuit(context.Context, string) (*entities.User, error)
	UpdateUser(context.Context, *entities.User) error
	SoftDeleteUser(context.Context, string) error
}
