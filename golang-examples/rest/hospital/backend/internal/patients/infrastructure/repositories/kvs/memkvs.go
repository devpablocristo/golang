package memkvsrepo

import (
	"encoding/json"
	"errors"

	"github.com/devpablocristo/golang/hex-arch/backend/internal/patients/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS(db map[string][]byte) *memkvs {
	return &memkvs{
		kvs: db,
	}
}

func (repo *memkvs) GetPatient(id string) (domain.Patient, error) {
	value, ok := repo.kvs[id]
	if ok {
		patient := domain.Patient{}
		err := json.Unmarshal(value, &patient)
		if err != nil {
			return domain.Patient{}, errors.New("fail to get value from kvs")
		}

		return patient, nil
	}

	return domain.Patient{}, errors.New("patient not found in kvs")
}

func (repo *memkvs) SavePatient(patient domain.Patient) error {
	newPatient, err := json.Marshal(patient)
	if err != nil {
		return errors.New("patient fails at marshal into json string")
	}

	repo.kvs[patient.Patient.UUID] = newPatient

	return nil
}
