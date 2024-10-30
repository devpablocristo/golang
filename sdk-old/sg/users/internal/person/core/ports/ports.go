package ports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
)

type UseCases interface {
	CreatePerson(ctx context.Context, person *entities.Person) (string, error)
	FindPersonByCuil(context.Context, string) (*entities.Person, error)
	FindPersonByUUID(context.Context, string) (*entities.Person, error)
	UpdatePersonByCuil(context.Context, *entities.Person) (string, error)
}

type Repository interface {
	CreatePerson(context.Context, *entities.Person) error
	UpdatePerson(context.Context, *entities.Person) error
	FindPersonByCuil(context.Context, string) (*entities.Person, error)
	FindPersonByUUID(context.Context, string) (*entities.Person, error)
}
