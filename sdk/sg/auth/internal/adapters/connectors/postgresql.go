package authconn

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"

	entities "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

const dbNameKey = "AUTH_DB"

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgreSQL() (ports.Repository, error) {
	fmt.Println(viper.GetString(dbNameKey))
	r, err := sdkpg.Bootstrap(dbNameKey)
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &PostgreSQL{
		repository: r,
	}, nil
}

func (s *PostgreSQL) CreateEvent(ctx context.Context) (*entities.Auth, error) {
	return nil, nil
}