package authconn

import (
	"context"

	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"

	entities "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
)

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgresSQL(r sdkpgports.Repository) ports.Repository {
	return &PostgreSQL{
		repository: r,
	}
}

func (s *PostgreSQL) CreateEvent(ctx context.Context) (*entities.Auth, error) {
	return nil, nil
}
