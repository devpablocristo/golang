package patientservice

import (
	"errors"

	domain "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/domain"
	ports "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/service/ports"
)

type PatientService struct {
	patientRepository ports.Repository
}

// func New(patientRepository ports.PatientRepository, uidGen uidgen.UIDGen) *service {
func NewPatientService(patientRepository ports.Repository) *PatientService {
	return &PatientService{
		patientRepository: patientRepository,
	}
}

func (ps *PatientService) GetPatient(id string) (domain.Patient, error) {
	patient, err := ps.patientRepository.GetPatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("get patient from repository has failed")
	}

	return patient, nil
}

func (ps *PatientService) CreatePatient(patient domain.Patient) (domain.Patient, error) {
	err := ps.patientRepository.SavePatient(patient)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

type ResponseInfo struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}
