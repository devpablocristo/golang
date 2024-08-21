package person

import (
	"context"
)

type MapDB struct {
	mDB map[string]*Person
}

func NewMapDB() *MapDB {
	m := make(map[string]*Person)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) SavePerson(ctx context.Context, p *Person) error {
	m.mDB[p.UUID] = p
	return nil
}

func (m *MapDB) GetPerson(ctx context.Context, UUID string) (*Person, error) {
	p := m.mDB[UUID]
	return p, nil
}

func (m *MapDB) ListPersons(ctx context.Context) map[string]*Person {
	// var results []Person
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
