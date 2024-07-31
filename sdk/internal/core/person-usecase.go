package core

import (
	"context"

	"github.com/devpablocristo/qh/events/internal/core/person"
)

type PersonUseCasePort interface {
	CreatePerson(context.Context, *person.Person) error
	// DeletePerson(context.Context, string) error
	// HardDeletePerson(context.Context, string) (person.Person, error)
	// UpdatePerson(context.Context, person.Person, string) (person.Person, error)
	// RevivePerson(context.Context, string) (person.Person, error)
	// GetPerson(context.Context, string) (*person.Person, error)
	// ListPersons(context.Context) ([]person.Person, error)
}

type PersonUseCase struct {
	storage person.RepositoryPort
}

func NewPersonService(s person.RepositoryPort) PersonUseCasePort {
	return &PersonUseCase{
		storage: s,
	}
}

func (ps *PersonUseCase) CreatePerson(ctx context.Context, p *person.Person) error {
	if err := ps.storage.SavePerson(ctx, p); err != nil {
		return err
	}
	return nil
}

// func (ps *PersonUseCase) GetPersons(ctx context.Context) (map[string]*person.Person, error) {
// 	persons := ps.storage.ListPersons(ctx)
// 	return persons, nil
// }

// func (ps *PersonUseCase) GetPerson(ctx context.Context, UUID string) (*person.Person, error) {
// 	p, err := ps.storage.GetPerson(ctx, UUID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }

// func (ps *PersonUseCase) UpdatePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.UpdatePerson(ctx, UUID)
// }

// func (ps *PersonUseCase) DeletePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.DeletePerson(ctx, UUID)
// }
