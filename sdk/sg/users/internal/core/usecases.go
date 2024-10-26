package user

import (
	"context"

	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
	"github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
	userports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
	personports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
	"github.com/google/uuid"
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
	personUUID, err := u.personUseCases.CreatePerson(ctx, dto.ToPerson(userDto.Person))
	if err != nil {
		return "", err
	}

	user := dto.ToUser(userDto)
	user.UUID = uuid.New().String()
	user.PersonUUID = &personUUID
	userModel, err := entities.ToDataModel(user)
	if err != nil {
		return "", err
	}
	if err := u.repository.CreateUser(ctx, userModel); err != nil {
		return "", err
	}

	return user.UUID, nil
}
