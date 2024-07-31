package person

import (
	"context"
)

type RepositoryPort interface {
	SavePerson(context.Context, *Person) error
	// DeletePerson(context.Context, string) (*Person, error)
	// HardDeletePerson(context.Context, string) (*Person, error)
	// UpdatePerson(context.Context, *Person, string) (*Person, error)
	// RevivePerson(context.Context, string) (*Person, error)
	// GetPerson(context.Context, string) (*Person, error)
	// ListPersons(context.Context) ([]Person, error)
}

type DAOPort interface {
	Create(context.Context, *Person) error
	// FindByID(context.Context, string) (*Person, error)
	// Update(context.Context, *Person, string) (*Person, error)
	// HardDelete(context.Context, string) (*Person, error)
	// List(context.Context) ([]Person, error)
	// SoftDelete(context.Context, string) (*Person, error)
	// SoftUndelete(context.Context, string) (*Person, error)
}
