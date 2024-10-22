package ports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
)

type Repository interface {
	Create(context.Context, *entities.Person) error
	FindByID(context.Context, string) (*entities.Person, error)
	Update(context.Context, *entities.Person) error
	SoftDelete(context.Context, string) error
	FindByCuit(context.Context, string) (*entities.Person, error)
}
