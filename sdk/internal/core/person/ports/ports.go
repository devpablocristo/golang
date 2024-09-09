package ports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/person/entities"
)

type UseCases interface {
	CreatePerson(context.Context, *entities.Person) error
	// DeletePerson(context.Context, string) error
	// HardDeletePerson(context.Context, string) (Person, error)
	// UpdatePerson(context.Context, Person, string) (Person, error)
	// RevivePerson(context.Context, string) (Person, error)
	// GetPerson(context.Context, string) (*entities.Person, error)
}

type Repository interface {
	SavePerson(context.Context, *entities.Person) error
	// DeletePerson(context.Context, string) (*entities.Person, error)
	// HardDeletePerson(context.Context, string) (*entities.Person, error)
	// UpdatePerson(context.Context, *entities.Person, string) (*entities.Person, error)
	// RevivePerson(context.Context, string) (*entities.Person, error)
	// GetPerson(context.Context, string) (*entities.Person, error)
	// ListPersons(context.Context) ([]Person, error)
}
