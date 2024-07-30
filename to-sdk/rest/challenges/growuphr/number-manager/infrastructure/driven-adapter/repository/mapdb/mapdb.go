package mapdb

import (
	"context"
	"errors"
	"sync"

	"github.com/devpablocristo/growuphr/number-manager/domain"
)

var (
	once  sync.Once
	newDB *MapDB
)

type MapDB struct {
	mDB map[string]*domain.ReservedNumber
	mux sync.Mutex
}

func NewMapDB() *MapDB {
	once.Do(func() {
		m := make(map[string]*domain.ReservedNumber)
		newDB = &MapDB{
			mDB: m,
			mux: sync.Mutex{},
		}
	})
	return newDB
}

func (m *MapDB) Create(ctx context.Context, rn *domain.ReservedNumber) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.mDB[rn.UUID] = rn
	_, exist := m.mDB[rn.UUID]
	if !exist {
		return errors.New("value not found")
	}

	return nil
}

func (m *MapDB) Read(ctx context.Context, UUID string) (*domain.ReservedNumber, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	rn, exist := m.mDB[UUID]
	if !exist {
		return nil, errors.New("value not found")
	}
	return rn, nil
}

func (m *MapDB) CheckForUsername(ctx context.Context, checkUsr string) (*domain.ReservedNumber, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()

	for _, rn := range m.mDB {
		if rn.User.Username == checkUsr {
			return rn, true
		}
	}
	return nil, false
}

func (m *MapDB) CheckForNumber(ctx context.Context, checkNum int) (*domain.ReservedNumber, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()

	for _, rn := range m.mDB {
		if rn.Number.Number == checkNum {
			return rn, true
		}
	}
	return nil, false
}

func (m *MapDB) List(ctx context.Context) map[string]*domain.ReservedNumber {
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
