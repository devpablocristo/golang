package ports

import (
	"context"

	datmod "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors/data-model"
	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UseCases interface {
	CreateUser(context.Context, *dto.UserDto) (string, error)
	UpdateUserByPersonCuil(context.Context, *dto.UserDto) (string, error)
	FindUserByPersonCuil(context.Context, string) (*entities.User, error)
	FindUserByPersonUUID(context.Context, string) (*entities.User, error)
	FindUserByUserUUID(context.Context, string) (*entities.User, error)
}

type Repository interface {
	CreateUser(context.Context, *datmod.User) error
	FindUserByPersonUUID(context.Context, string) (*entities.User, error)
	FindUserByUserUUID(context.Context, string) (*entities.User, error)
	UpdateUser(context.Context, *datmod.User) error
}
