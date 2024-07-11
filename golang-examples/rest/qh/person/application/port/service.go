package port

import (
	"context"

	"github.com/devpablocristo/golang/06-projects/qh/person/domain"
)

type Service interface {
	GetPersons(context.Context) (map[string]*domain.Person, error)
	GetPerson(context.Context, string) (*domain.Person, error)
	CreatePerson(context.Context, *domain.Person) (*domain.Person, error)
	UpdatePerson(context.Context, string) error
	DeletePerson(context.Context, string) error
}
