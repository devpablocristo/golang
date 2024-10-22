package user

import (
	"context"

	companyports "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/ports"
	userports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
	personports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

type useCases struct {
	userRepo    userports.Repository
	personRepo  personports.Repository
	companyRepo companyports.Repository
}

func NewUseCases(ur userports.Repository, pr personports.Repository, cr companyports.Repository) userports.UseCases {
	return &useCases{
		userRepo:    ur,
		personRepo:  pr,
		companyRepo: cr,
	}
}

func (u *useCases) CheckUserStatus(ctx context.Context, cuit string) (bool, error) {
	userFound, err := u.findUserByCuit(ctx, cuit)
	if err != nil {
		return false, err
	}
	if userFound {
		return true, nil
	}
	adminRequestFound, err := u.findAdministrativeRequestByCuit(ctx, cuit)
	if err != nil {
		return false, err
	}
	if adminRequestFound {
		return true, nil
	}

	return false, nil
}
