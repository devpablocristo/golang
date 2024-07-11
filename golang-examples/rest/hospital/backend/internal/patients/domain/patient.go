package domain

import (
	doctor "github.com/devpablocristo/golang/hex-arch/backend/internal/doctors/domain"
	person "github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
)

type Patient struct {
	Patient   person.Person `json:"patient"`
	Doctor    doctor.Doctor `json:"doctor"`
	Hospital  string        `json:"hospital"`
	Diagnosis string        `json:"diagnosis"`
}
