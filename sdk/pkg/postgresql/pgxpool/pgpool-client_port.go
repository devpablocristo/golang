package pgxpostgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQLClientPort interface {
	Connect(config PostgreSQLClientConfig) error
	Close()
	Pool() *pgxpool.Pool
}
