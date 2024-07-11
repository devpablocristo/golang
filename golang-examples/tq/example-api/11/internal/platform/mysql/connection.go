package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBHost = "mysql-local" // Docker container name
	DBPort = 3306
	DBName = "meli_items"
	DBUser = "root"
	DBPass = "root"
)

var db *sqlx.DB //nolint:gochecknoglobals

// GetConnectionDB establishes a connection to the database and returns it.
func GetConnectionDB() (*sqlx.DB, error) {
	var err error

	if db == nil {
		db, err = sqlx.Connect("mysql", dbConnectionURL())
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}
	}

	return db, nil
}

// dbConnectionURL builds the MySQL connection string.
func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		DBUser, DBPass, DBHost, DBPort, DBName)
}
