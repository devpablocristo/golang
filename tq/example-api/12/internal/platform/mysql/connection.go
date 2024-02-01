package mysql

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBHost = "127.0.0.1" // Use docker container name or localhost
	DBPort = 3306
	DBName = "meli_items" // Adjust according to your needs
	DBUser = "root"
	DBPass = "root"
)

var db *sqlx.DB // Use gochecknoglobals for ignoring global variable warning if necessary

// GetConnectionDB establishes a connection to the database and returns it.
// It also ensures that the database is up to date by running migrations.
func GetConnectionDB() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	db, err = sqlx.Connect("mysql", dbConnectionURL())
	if err != nil {
		log.Printf("########## DB ERROR: %s #############", err)
		return nil, fmt.Errorf("DB ERROR: %w", err)
	}

	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// autoMigrate applies database migrations automatically.
func autoMigrate(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		log.Printf("########## DB ERROR: %s #############", err)
		return fmt.Errorf("error instantiating migration: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://path/to/your/migrations",
		"mysql", driver,
	)
	if err != nil {
		log.Printf("########## DB ERROR: %s #############", err)
		return fmt.Errorf("error instantiating migration: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error executing migration: %w", err)
	}

	return nil
}

// dbConnectionURL builds the MySQL connection string.
func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName)
}
