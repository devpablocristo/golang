package person

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
	"github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
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
	if err := u.repository.CreatePerson(ctx, person); err != nil {
		return "", fmt.Errorf("failed to create person: %w", err)
	}

	return person.UUID, nil
}

func (u *useCases) FindPersonByCuil(ctx context.Context, cuil string) (*entities.Person, error) {
	// Buscar persona existente por CUIL usando el repositorio
	person, err := u.repository.FindPersonByCuil(ctx, cuil)
	if err != nil {
		return nil, fmt.Errorf("failed to find person by CUIL: %w", err)
	}

	return person, nil
}

func (u *useCases) UpdatePersonByCuil(ctx context.Context, person *entities.Person) (string, error) {
	// Buscar persona existente por CUIL
	existingPerson, err := u.FindPersonByCuil(ctx, person.Cuil)
	if err != nil {
		return "", fmt.Errorf("failed to find person by CUIL: %w", err)
	}

	// Flag para determinar si hubo cambios
	hasChanges := false

	// Actualizar solo los campos que no estén vacíos o que hayan cambiado
	if person.Dni != "" && person.Dni != existingPerson.Dni {
		existingPerson.Dni = person.Dni
		hasChanges = true
	}
	if person.FirstName != "" && person.FirstName != existingPerson.FirstName {
		existingPerson.FirstName = person.FirstName
		hasChanges = true
	}
	if person.LastName != "" && person.LastName != existingPerson.LastName {
		existingPerson.LastName = person.LastName
		hasChanges = true
	}
	if person.Nationality != "" && person.Nationality != existingPerson.Nationality {
		existingPerson.Nationality = person.Nationality
		hasChanges = true
	}
	if person.Email != "" && person.Email != existingPerson.Email {
		existingPerson.Email = person.Email
		hasChanges = true
	}
	if person.Phone != "" && person.Phone != existingPerson.Phone {
		existingPerson.Phone = person.Phone
		hasChanges = true
	}

	// Actualizar el campo UpdatedAt y guardar cambios solo si hubo modificaciones
	if hasChanges {

		// Guardar los cambios en el repositorio
		if err := u.repository.UpdatePerson(ctx, existingPerson); err != nil {
			return "", fmt.Errorf("failed to update person: %w", err)
		}
	}

	return existingPerson.UUID, nil
}

func (u *useCases) FindPersonByUUID(ctx context.Context, uuid string) (*entities.Person, error) {
	// Buscar persona existente por UUID usando el repositorio
	person, err := u.repository.FindPersonByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find person by UUID: %w", err)
	}
	return person, nil
}
