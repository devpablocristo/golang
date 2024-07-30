package patients

import (
	"database/sql"

	domain "github.com/devpablocristo/interviews/hospital2/patients/src/domain"
)

type PatientGateway interface {
	CreatePatient(p *domain.CreatePatientCMD) (*domain.Patient, error)
	GetPatients() []*domain.Patient
	GetPatientByID(id int64) (*domain.Patient, error)
}
type CreatePatientInDB struct {
	PatientStorage
}

func (c *CreatePatientInDB) CreatePatient(p *domain.CreatePatientCMD) (*domain.Patient, error) {
	return c.createPatientDB(p)
}

func (c *CreatePatientInDB) GetPatients() []*domain.Patient {
	return c.getPatientsDB()
}

func (c *CreatePatientInDB) GetPatientByID(id int64) (*domain.Patient, error) {
	return c.getPatientByIDBD(id)
}

func NewPatientGateway(db *sql.DB) PatientGateway {
	return &CreatePatientInDB{NewPatientStorageGateway(db)}
}
