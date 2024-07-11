package inmemorydb

import (
	"fmt"

	"github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
)

type InmemoryDB struct {
	Storage map[string]domain.Person
}

func NewInmemoryDB(s map[string]domain.Person) *InmemoryDB {
	return &InmemoryDB{
		Storage: s,
	}
}

func (s *InmemoryDB) Exists(uuid string) bool {
	_, found := s.Storage[uuid]
	return found
}

func (s *InmemoryDB) Add(p domain.Person) {

	fmt.Println(p)
	s.Storage[p.UUID] = p
}

func (s *InmemoryDB) List() map[string]domain.Person {
	return s.Storage
}
