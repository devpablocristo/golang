package mapdb

import (
	"context"
	"errors"
	"sync"

	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

var (
	once  sync.Once
	newDB *MapDB
)

type MapDB struct {
	mDB map[string]*domain.User
	mux sync.Mutex
}

func NewMapDB() *MapDB {
	once.Do(func() {
		m := make(map[string]*domain.User)
		newDB = &MapDB{
			mDB: m,
			mux: sync.Mutex{},
		}
	})
	return newDB
}

func (m *MapDB) Create(ctx context.Context, us *domain.User) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.mDB[us.UUID] = us
	_, exist := m.mDB[us.UUID]
	if !exist {
		return errors.New("user not found")
	}

	return nil
}

func (m *MapDB) Read(ctx context.Context, UUID string) (*domain.User, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	or, exist := m.mDB[UUID]
	if !exist {
		return nil, errors.New("user not found")
	}
	return or, nil
}

func (m *MapDB) List(ctx context.Context) map[string]*domain.User {
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

func (m *MapDB) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	for _, user := range m.mDB {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
