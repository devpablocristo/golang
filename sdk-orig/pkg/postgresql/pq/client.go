package pqpostgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Singleton instance and related variables
var (
	instance PostgreSQLClientPort
	once     sync.Once
	errInit  error
	db       *sql.DB
)

type PostgreSQLClientPort interface {
	Connect(config PostgreSQLClientConfig) error
	Close()
	DB() *sql.DB
}


// PostgreSQLClient struct
type PostgreSQLClient struct {
	db *sql.DB
}

func InitializePostgreSQLClient(config PostgreSQLClientConfig) error {
	once.Do(func() {
		instance = &PostgreSQLClient{}
		errInit = instance.Connect(config)
		if errInit != nil {
			instance = nil
		}
	})
	return errInit
}

func GetPostgreSQLInstance() (PostgreSQLClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("PostgreSQL client is not initialized")
	}
	return instance, nil
}

func (client *PostgreSQLClient) Connect(config PostgreSQLClientConfig) error {
	connString := BuildConnString(config)
	var err error
	client.db, err = sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	err = client.db.Ping()
	if err != nil {
		return fmt.Errorf("unable to ping the database: %w", err)
	}

	return nil
}

func (client *PostgreSQLClient) Close() {
	if client.db != nil {
		client.db.Close()
	}
}

func (client *PostgreSQLClient) DB() *sql.DB {
	return client.db
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}

func ApplyMigrations(db *sql.DB, dbName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///app/migrations", dbName, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}