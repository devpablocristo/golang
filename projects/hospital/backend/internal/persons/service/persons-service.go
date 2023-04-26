package application

import (
	"errors"

	"github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
	"github.com/devpablocristo/golang/hex-arch/backend/internal/persons/service/ports"
)

var (
	ErrPersonExists = errors.New("person exits")
)

type PersonService struct {
	storage ports.Storage
}

func NewPersonaApplication(s ports.Storage) *PersonService {
	return &PersonService{
		storage: s,
	}
}

func (ps *PersonService) CreatePerson(p domain.Person) error {
	if ps.storage.Exists(p.UUID) {
		return ErrPersonExists
	}

	ps.storage.Add(p)

	return nil
}

func (ps *PersonService) List() map[string]domain.Person {
	return ps.storage.List()
}
