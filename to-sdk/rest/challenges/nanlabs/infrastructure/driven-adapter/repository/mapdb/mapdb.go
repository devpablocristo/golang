package mapdb

import (
	"context"

	domain "github.com/devpablocristo/nanlabs/domain"
)

type MapDB struct {
	mDB map[string]*domain.Task
}

func NewMapDB() *MapDB {
	m := make(map[string]*domain.Task)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) SaveTask(ctx context.Context, p *domain.Task) error {
	m.mDB[p.UUID] = p
	return nil
}

func (m *MapDB) GetTask(ctx context.Context, UUID string) (*domain.Task, error) {
	p := m.mDB[UUID]
	return p, nil
}

func (m *MapDB) ListTasks(ctx context.Context) map[string]*domain.Task {
	return m.mDB
}

func (m *MapDB) DeleteTask(ctx context.Context, UUID string) error {
	delete(m.mDB, UUID)
	return nil
}

func (m *MapDB) UpdateTask(ctx context.Context, UUID string) error {
	return nil
}
