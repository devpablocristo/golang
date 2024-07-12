package mysql

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
)

// Environment variables for database configuration
var (
	DBHost = getEnv("DB_HOST", "127.0.0.1")  // Default to "127.0.0.1" if not set
	DBPort = getEnv("DB_PORT", "3306")       // Default to "3306" if not set
	DBName = getEnv("DB_NAME", "meli_items") // Default to "meli_items" if not set
	DBUser = getEnv("DB_USER", "root")       // Default to "root" if not set
	DBPass = getEnv("DB_PASS", "root")       // Default to "root" if not set
)

// getEnv retrieves environment variables or returns a default value if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

var db *sqlx.DB // Global variable for the database connection

// GetConnectionDB establishes and returns a database connection
func GetConnectionDB() (*sqlx.DB, error) {
	if db != nil {
		return db, nil // Return the existing connection if already established
	}

	dsn := dbConnectionURL() // Build the data source name
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Printf("DB Connection Error: %s", err)
		return nil, fmt.Errorf("DB Connection Error: %w", err)
	}

	if err := autoMigrate(); err != nil {
		return nil, err
	}

	return db, nil
}

// autoMigrate applies database migrations automatically
func autoMigrate() error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		log.Printf("Migration Driver Error: %s", err)
		return fmt.Errorf("migration Driver Error: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://path/to/your/migrations", // Adjust this path
		"mysql", driver,
	)
	if err != nil {
		log.Printf("Migration Instance Error: %s", err)
		return fmt.Errorf("migration Instance Error: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration Error: %s", err)
		return fmt.Errorf("migration Error: %w", err)
	}

	return nil
}

// dbConnectionURL builds the MySQL connection string.
func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&multiStatements=true",
		DBUser, DBPass, DBHost, DBPort, DBName)
}
