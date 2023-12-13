package Person

import (
	"context"

	domain "github.com/devpablocristo/qh/internal/users/persons/domain"
)

type Repo interface {
	CreatePerson(context.Context, *domain.Person) (*domain.Person, error)
	// DeletePerson(context.Context, string) (*domain.Person, error)
	// HardDeletePerson(context.Context, string) (*domain.Person, error)
	// UpdatePerson(context.Context, *domain.Person, string) (*domain.Person, error)
	// RevivePerson(context.Context, string) (*domain.Person, error)
	// GetPerson(context.Context, string) (*domain.Person, error)
	// GetAllPersons(context.Context) ([]domain.Person, error)
}

type Service interface {
	CreatePerson(context.Context, *domain.Person) (*domain.Person, error)
	// DeletePerson(context.Context, string) (*domain.Person, error)
	// HardDeletePerson(context.Context, string) (*domain.Person, error)
	// UpdatePerson(context.Context, *domain.Person, string) (*domain.Person, error)
	// RevivePerson(context.Context, string) (*domain.Person, error)
	// GetPerson(context.Context, string) (*domain.Person, error)
	// GetAllPersons(context.Context) ([]domain.Person, error)
}
