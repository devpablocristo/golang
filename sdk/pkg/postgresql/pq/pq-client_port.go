package pqpostgresql

import "database/sql"

// PostgreSQLClientPort interface
type PostgreSQLClientPort interface {
	Connect(config PostgreSQLClientConfig) error
	Close()
	DB() *sql.DB
}
