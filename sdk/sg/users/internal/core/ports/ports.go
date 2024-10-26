package ports

import (
	"context"

	datmod "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors/data-model"
	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
	//entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UseCases interface {
	CreateUser(context.Context, *dto.UserDto) (string, error)
}

type Repository interface {
	CreateUser(context.Context, *datmod.User) error
	// FindUserByCuil(context.Context, string) (*entities.User, error)
	// UpdateUser(context.Context, *datmod.User) error
}
