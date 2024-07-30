package handler

import "github.com/devpablocristo/qh/internal/users/persons/domain"

type PersonDTO struct {
	Name          string
	Email         string
	Qualification int `validate:"gte=1,lte=10"`
}

func dtoToDomain(u *PersonDTO) *domain.Person {
	return &domain.Person{
		Name:          u.Name,
		Email:         u.Email,
		Qualification: u.Qualification,
	}
}
