package domain

import (
	person "github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
)

type Doctor struct {
	Doctor     person.Person `json:"uuid"`
	Speciality string        `json:"speciality"`
}
