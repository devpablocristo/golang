package person

import (
	"context"
)

type RepoPort interface {
	CreatePerson(context.Context, *Person) (*Person, error)
	// DeletePerson(context.Context, string) (Person, error)
	// HardDeletePerson(context.Context, string) (Person, error)
	// UpdatePerson(context.Context, Person, string) (Person, error)
	// RevivePerson(context.Context, string) (Person, error)
	// GetPerson(context.Context, string) (Person, error)
	// GetAllPersons(context.Context) ([]domain.Person, error)
}

type UseCasePort interface {
	CreatePerson(context.Context, Person) (Person, error)
	// DeletePerson(context.Context, string) (Person, error)
	// HardDeletePerson(context.Context, string) (Person, error)
	// UpdatePerson(context.Context, Person, string) (Person, error)
	// RevivePerson(context.Context, string) (Person, error)
	// GetPerson(context.Context, string) (Person, error)
	// GetAllPersons(context.Context) ([]domain.Person, error)
}
