package person

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/person/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/person/ports"
)

type useCases struct {
	storage ports.Repository
}

func NewPersonService(s ports.Repository) ports.UseCases {
	return &useCases{
		storage: s,
	}
}

func (ps *useCases) CreatePerson(ctx context.Context, p *entities.Person) error {
	if err := ps.storage.SavePerson(ctx, p); err != nil {
		return err
	}
	return nil
}

// func (ps *useCases) GetPersons(ctx context.Context) (map[string]*entities.Person, error) {
// 	persons := ps.storage.ListPersons(ctx)
// 	return persons, nil
// }

// func (ps *useCases) GetPerson(ctx context.Context, UUID string) (*entities.Person, error) {
// 	p, err := ps.storage.GetPerson(ctx, UUID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }

// func (ps *useCases) UpdatePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.UpdatePerson(ctx, UUID)
// }

// func (ps *useCases) DeletePerson(ctx context.Context, UUID string) error {
// 	return ps.storage.DeletePerson(ctx, UUID)
// }
