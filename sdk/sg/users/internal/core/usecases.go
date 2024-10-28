package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"
	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
	userports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
	personports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

type useCases struct {
	repository     userports.Repository
	personUseCases personports.UseCases
}

func NewUseCases(r userports.Repository, pu personports.UseCases) userports.UseCases {
	return &useCases{
		repository:     r,
		personUseCases: pu,
	}
}

func (u *useCases) CreateUser(ctx context.Context, userDto *dto.UserDto) (string, error) {
	// Convertir UserDto a entidad Person y crearla
	person := dto.ToPerson(userDto.Person)
	personUUID, err := u.personUseCases.CreatePerson(ctx, person)
	if err != nil {
		return "", fmt.Errorf("failed to create person: %w", err)
	}

	// Convertir UserDto a entidad User
	user := dto.ToUser(userDto)
	user.UUID = uuid.New().String()
	user.PersonUUID = &personUUID

	// Hashear la contraseña
	hashedPassword, err := sdktools.HashPassword(user.Credentials.Password, 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	user.Credentials.Password = hashedPassword

	// Convertir User a modelo de datos y crear el usuario
	userModel, err := entities.ToDataModel(user)
	if err != nil {
		return "", fmt.Errorf("failed to convert user to data model: %w", err)
	}
	if err := u.repository.CreateUser(ctx, userModel); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return user.UUID, nil
}

func (u *useCases) UpdateUserByPersonCuil(ctx context.Context, userDto *dto.UserDto) (string, error) {
	// Actualizar la persona asociada al usuario por CUIL
	person := dto.ToPerson(userDto.Person)
	personUUID, err := u.personUseCases.UpdatePersonByCuil(ctx, person)
	if err != nil {
		return "", fmt.Errorf("failed to update person by CUIL: %w", err)
	}

	// Buscar usuario existente por PersonUUID
	existingUser, err := u.repository.FindUserByPersonUUID(ctx, personUUID)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	// Flag para determinar si hubo cambios
	hasChanges := false

	// Actualizar campos si tienen nuevos valores
	if userDto.EmailValidated != existingUser.EmailValidated {
		existingUser.EmailValidated = userDto.EmailValidated
		hasChanges = true
	}

	if userDto.Credentials.Username != "" && userDto.Credentials.Username != existingUser.Credentials.Username {
		existingUser.Credentials.Username = userDto.Credentials.Username
		hasChanges = true
	}

	if userDto.Credentials.Password != "" {
		// Hashear la nueva contraseña
		hashedPassword, err := sdktools.HashPassword(userDto.Credentials.Password, 12)
		if err != nil {
			return "", fmt.Errorf("failed to hash password: %w", err)
		}
		existingUser.Credentials.Password = hashedPassword
		hasChanges = true
	}

	// Actualizar roles si son diferentes
	if len(userDto.Roles) > 0 {
		newRoles := mapRoles(userDto.Roles)
		if !rolesAreEqual(existingUser.Roles, newRoles) {
			existingUser.Roles = newRoles
			hasChanges = true
		}
	}

	// Actualizar el campo UpdatedAt y guardar cambios si hubo modificaciones
	if hasChanges {
		now := time.Now()
		existingUser.UpdatedAt = &now

		userModel, err := entities.ToDataModel(existingUser)
		if err != nil {
			return "", fmt.Errorf("failed to convert user to data model: %w", err)
		}

		if err := u.repository.UpdateUser(ctx, userModel); err != nil {
			return "", fmt.Errorf("failed to update user: %w", err)
		}
	}

	return existingUser.UUID, nil
}

func (u *useCases) FindUserByPersonCuil(ctx context.Context, cuil string) (*entities.User, error) {
	// Step 1: Find the person by CUIL using personUseCases
	person, err := u.personUseCases.FindPersonByCuil(ctx, cuil)
	if err != nil {
		return nil, fmt.Errorf("failed to find person by CUIL: %w", err)
	}
	if person == nil {
		return nil, fmt.Errorf("person with CUIL %s not found", cuil)
	}

	// Step 2: Find the user by the person's UUID
	user, err := u.repository.FindUserByPersonUUID(ctx, person.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by PersonUUID: %w", err)
	}

	return user, nil
}

// Implementación de FindUserByPersonUUID
func (u *useCases) FindUserByPersonUUID(ctx context.Context, personUUID string) (*entities.User, error) {
	user, err := u.repository.FindUserByPersonUUID(ctx, personUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by PersonUUID: %w", err)
	}
	return user, nil
}

// Implementación de FindUserByUserUUID
func (u *useCases) FindUserByUserUUID(ctx context.Context, userUUID string) (*entities.User, error) {
	user, err := u.repository.FindUserByUserUUID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by UserUUID: %w", err)
	}
	return user, nil
}
