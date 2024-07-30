package mapdb

import (
	"context"
	"errors"
	"sync"

	"github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

var (
	once  sync.Once
	newDB *MapDB
)

type MapDB struct {
	mDB map[string]*domain.Order
	mux sync.Mutex
}

func NewMapDB() *MapDB {
	once.Do(func() {
		m := make(map[string]*domain.Order)
		newDB = &MapDB{
			mDB: m,
			mux: sync.Mutex{},
		}
	})
	return newDB
}

func (m *MapDB) Create(ctx context.Context, or *domain.Order) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.mDB[or.UUID] = or
	_, exist := m.mDB[or.UUID]
	if !exist {
		return errors.New("order not found")
	}

	return nil
}

func (m *MapDB) Read(ctx context.Context, UUID string) (*domain.Order, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	or, exist := m.mDB[UUID]
	if !exist {
		return nil, errors.New("order not found")
	}
	return or, nil
}

func (m *MapDB) List(ctx context.Context) map[string]*domain.Order {
	m.mux.Lock()
	defer m.mux.Unlock()

	return m.mDB
}

func (m *MapDB) Delete(ctx context.Context, UUID string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.mDB, UUID)
	return nil
}
