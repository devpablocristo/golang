package ports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UseCases interface {
	CheckCuit(context.Context, string) (bool, error)
}

type Repository interface {
	Create(context.Context, *entities.User) error
	FindByUUID(context.Context, string) (*entities.User, error)
	FindByCuit(context.Context, string) (*entities.User, error)
	Update(context.Context, *entities.User) error
	SoftDelete(context.Context, string) error
}
