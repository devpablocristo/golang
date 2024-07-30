package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go-micro.dev/v4/logger"

	cfg "github.com/devpablocristo/qh/events/internal/platform/config"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

var (
	instance *PostgreSQL
	once     sync.Once
)

func NewPostgreSQL(dbConfig cfg.DBConfig) *PostgreSQL {
	once.Do(func() {
		instance = &PostgreSQL{}
		if err := instance.init(dbConfig); err != nil {
			logger.Fatalf("Failed to initialize PostgreSQL database: %v", err)
		}
	})
	return instance
}

func (s *PostgreSQL) init(dbConfig cfg.DBConfig) error {
	connString := s.buildConnString(dbConfig)

	db, err := s.openDBConnection(connString)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := s.checkDBExists(db, dbConfig.DBName); err != nil {
		return err
	}

	if err := s.applyMigrations(db, dbConfig.DBName); err != nil {
		return err
	}

	return s.connectPool(connString)
}

func (s *PostgreSQL) buildConnString(dbConfig cfg.DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DBName)
}

func (s *PostgreSQL) openDBConnection(connString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func (s *PostgreSQL) checkDBExists(db *sql.DB, dbName string) error {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", dbName)
	if err := db.QueryRow(query).Scan(&exists); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		return fmt.Errorf("database %s does not exist", dbName)
	}
	return nil
}

func (s *PostgreSQL) applyMigrations(db *sql.DB, dbName string) error {
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

func (s *PostgreSQL) connectPool(connString string) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return fmt.Errorf("unable to parse database connection string: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping the database: %w", err)
	}

	log.Println("Connected to the PostgreSQL database successfully")
	s.pool = pool
	return nil
}

func (s *PostgreSQL) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

func (s *PostgreSQL) Pool() *pgxpool.Pool {
	return s.pool
}
