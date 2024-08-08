package core

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/person"
)

type PersonUseCasesPort interface {
	CreatePerson(context.Context, *person.Person) error
	// DeletePerson(context.Context, string) error
	// HardDeletePerson(context.Context, string) (person.Person, error)
	// UpdatePerson(context.Context, person.Person, string) (person.Person, error)
	// RevivePerson(context.Context, string) (person.Person, error)
	// GetPerson(context.Context, string) (*person.Person, error)
	// ListPersons(context.Context) ([]person.Person, error)
}

type PersonUseCases struct {
	storage person.RepositoryPort
}

func NewPersonService(s person.RepositoryPort) PersonUseCasesPort {
	return &PersonUseCases{
		storage: s,
	}
}

func (ps *PersonUseCases) CreatePerson(ctx context.Context, p *person.Person) error {
	if err := ps.storage.SavePerson(ctx, p); err != nil {
		return err
	}
	return nil
}

// func (ps *PersonUseCases) GetPersons(ctx context.Context) (map[string]*person.Person, error) {
// 	persons := ps.storage.ListPersons(ctx)
// 	return persons, nil
// }

// func (ps *PersonUseCases) GetPerson(ctx context.Context, UUID string) (*person.Person, error) {
// 	p, err := ps.storage.GetPerson(ctx, UUID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }

// func (ps *PersonUseCases) UpdatePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.UpdatePerson(ctx, UUID)
// }

// func (ps *PersonUseCases) DeletePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.DeletePerson(ctx, UUID)
// }
