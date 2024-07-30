package mapdb

import (
	"context"

	domain "github.com/devpablocristo/golang/06-projects/qh/person/domain"
)

type MapDB struct {
	mDB map[string]*domain.Person
}

func NewMapDB() *MapDB {
	m := make(map[string]*domain.Person)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) SavePerson(ctx context.Context, p *domain.Person) error {
	m.mDB[p.UUID] = p
	return nil
}

func (m *MapDB) GetPerson(ctx context.Context, UUID string) (*domain.Person, error) {
	p := m.mDB[UUID]
	return p, nil
}

func (m *MapDB) ListPersons(ctx context.Context) map[string]*domain.Person {
	// var results []domain.Person
	// for _, person := range m.mDB {
	// 	results = append(results, *person)
	// }
	// return results, nil

	return m.mDB
}

func (m *MapDB) DeletePerson(ctx context.Context, UUID string) error {
	delete(m.mDB, UUID)
	return nil
}

func (m *MapDB) UpdatePerson(ctx context.Context, UUID string) error {
	return nil
}
