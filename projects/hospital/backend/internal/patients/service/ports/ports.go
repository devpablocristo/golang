package ports

import patient "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/domain"

// patient http handler port
type Service interface {
	GetPatient(id string) (patient.Patient, error)
	CreatePatient(patient.Patient) (patient.Patient, error)
}

// patient repository port
//
//go:generate mockgen -source=./service.go -destination=../../../mocks/service_mock.go -package=mocks
type Repository interface {
	GetPatient(id string) (patient.Patient, error)
	SavePatient(patient.Patient) error
}

//go:generate mockgen -source=./service.go -destination=../../../mocks/service_mock.go -package=mocks
type Scrapper interface {
	Check() error
}
