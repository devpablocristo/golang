package port

import (
	"context"

	"github.com/devpablocristo/golang/06-projects/qh/person/domain"
)

type Storage interface {
	SavePerson(context.Context, *domain.Person) error
	GetPerson(context.Context, string) (*domain.Person, error)
	ListPersons(context.Context) map[string]*domain.Person
	DeletePerson(context.Context, string) error
	UpdatePerson(context.Context, string) error
}
