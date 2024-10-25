package user

import (
	"context"
	"time"

	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
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
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = nil
	if err := u.repository.CreateUser(ctx, user); err != nil {
		return "", err
	}

	return user.UUID, nil
}

// func (u *useCases) CheckUserStatus(ctx context.Context, cuil string) (bool, error) {
// 	userFound, err := u.FindUserByCuit(ctx, cuil)
// 	if err != nil {
// 		return false, err
// 	}
// 	if userFound {
// 		return true, nil
// 	}
// 	adminRequestFound, err := u.findAdministrativeRequestByCuit(ctx, cuil)
// 	if err != nil {
// 		return false, err
// 	}
// 	if adminRequestFound {
// 		return true, nil
// 	}

// 	return false, nil
// }
