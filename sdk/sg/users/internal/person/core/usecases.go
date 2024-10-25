package person

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
	ports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

type useCases struct {
	repository ports.Repository
}

func NewUseCases(r ports.Repository) ports.UseCases {
	return &useCases{
		repository: r,
	}
}

func (u *useCases) CreatePerson(ctx context.Context, person *entities.Person) (string, error) {
	// Generar UUID y establecer timestamps
	person.UUID = uuid.New().String()

	// Intentar crear la persona en el repositorio
	err := u.repository.CreatePerson(ctx, person)
	if err != nil {
		// Agregar contexto al error para debug
		return "", fmt.Errorf("failed to create person: %w", err)
	}

	return person.UUID, nil
}

func (u *useCases) FindPersonByCuil(context.Context, string) (*entities.Person, error) {
	return nil, nil
}
