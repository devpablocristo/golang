package gosqldriver

import "database/sql"

type MySQLClientPort interface {
	DB() *sql.DB
	Close()
}
