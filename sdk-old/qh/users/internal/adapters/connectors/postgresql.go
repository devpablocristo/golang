package userconn

import (
	"context"

	sdkports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
	entities "github.com/devpablocristo/golang/sdk/qh/users/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/qh/users/internal/core/ports"
)

type PostgreSQL struct {
	repository sdkports.Repository
}

func NewPostgresSQL(r sdkports.Repository) ports.Repository {
	return &PostgreSQL{
		repository: r,
	}
}

func (s *PostgreSQL) SaveUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (s *PostgreSQL) GetUserByUUID(ctx context.Context, UUID string) (*entities.User, error) {
	return nil, nil
}

func (s *PostgreSQL) GetUserByCredentials(ctx context.Context, username, passwordHash string) (string, error) {
	return "", nil
}

func (s *PostgreSQL) DeleteUser(ctx context.Context, UUID string) error {
	return nil
}

func (s *PostgreSQL) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	return nil, nil
}

func (s *PostgreSQL) UpdateUser(ctx context.Context, user *entities.User, UUID string) error {
	return nil
}
