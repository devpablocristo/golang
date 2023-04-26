package application

import (
	"context"

	"github.com/devpablocristo/golang/06-projects/qh/person/application/port"
	"github.com/devpablocristo/golang/06-projects/qh/person/domain"
)

type PersonService struct {
	storage port.Storage
}

func NewPersonService(s port.Storage) *PersonService {
	return &PersonService{
		storage: s,
	}
}

func (ps *PersonService) GetPersons(ctx context.Context) (map[string]*domain.Person, error) {
	persons := ps.storage.ListPersons(ctx)
	return persons, nil
}

func (ps *PersonService) GetPerson(ctx context.Context, UUID string) (*domain.Person, error) {
	p, err := ps.storage.GetPerson(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *PersonService) CreatePerson(ctx context.Context, p *domain.Person) (*domain.Person, error) {
	err := ps.storage.SavePerson(ctx, p)
	if err != nil {
		return &domain.Person{}, err
	}
	return p, nil
}

func (ps *PersonService) UpdatePerson(ctx context.Context, UUID string) error {
	return ps.storage.UpdatePerson(ctx, UUID)
}

func (ps *PersonService) DeletePerson(ctx context.Context, UUID string) error {
	return ps.storage.DeletePerson(ctx, UUID)
}
