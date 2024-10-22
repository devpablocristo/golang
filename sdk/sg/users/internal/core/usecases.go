package auth

import (
	"context"

	userports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
	personports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
	companyports "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/ports"

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

func (u *useCases) CheckCuit(ctx context.Context, cuit string) (bool, error) {
	// First, check if the CUIT belongs to a person
	person, err := u.personRepo.FindByCuit(ctx, cuit)
	if err != nil {
		return false, err
	}
	if person != nil {
		// CUIT already exists for a person
		return true, nil
	}

	// Next, check if the CUIT belongs to a company
	company, err := u.companyRepo.FindByCuit(ctx, cuit)
	if err != nil {
		return false, err
	}
	if company != nil {
		// CUIT already exists for a company
		return true, nil
	}

	// If not found in both, the CUIT does not exist
	return false, nil
}
