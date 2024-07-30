package pgxpostgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	instance PostgreSQLClientPort
	once     sync.Once
	errInit  error
)

type PostgreSQLClient struct {
	pool *pgxpool.Pool
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
	pool, err := ConnectPool(connString)
	if err != nil {
		return err
	}
	client.pool = pool
	return nil
}

func (client *PostgreSQLClient) Close() {
	if client.pool != nil {
		client.pool.Close()
	}
}

func (client *PostgreSQLClient) Pool() *pgxpool.Pool {
	return client.pool
}

func ConnectPool(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database connection string: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping the database: %w", err)
	}

	return pool, nil
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
