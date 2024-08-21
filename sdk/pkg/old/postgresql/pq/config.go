package pqpostgresql

import (
	"fmt"
)

type PostgreSQLClientConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func BuildConnString(config PostgreSQLClientConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DBName)
}

func (config PostgreSQLClientConfig) Validate() error {
	if config.User == "" {
		return fmt.Errorf("POSTGRES_USER is required")
	}
	if config.Password == "" {
		return fmt.Errorf("POSTGRES_PASSWORD is required")
	}
	if config.Host == "" {
		return fmt.Errorf("POSTGRES_HOST is required")
	}
	if config.Port == "" {
		return fmt.Errorf("POSTGRES_PORT is required")
	}
	if config.DBName == "" {
		return fmt.Errorf("POSTGRES_DB is required")
	}
	return nil
}