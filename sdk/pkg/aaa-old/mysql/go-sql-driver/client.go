package gosqldriver

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	instance  MySQLClientPort
	once      sync.Once
	initError error
)

type MySQLClientPort interface {
	DB() *sql.DB
	Close()
}

type MySQLClient struct {
	db *sql.DB
}

func InitializeMySQLClient(config MySQLClientConfig) error {
	once.Do(func() {
		client := &MySQLClient{}
		initError = client.connect(config)
		if initError != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return initError
}

func GetMySQLInstance() (MySQLClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("MySQL client is not initialized")
	}
	return instance, nil
}

func (client *MySQLClient) connect(config MySQLClientConfig) error {
	dsn := config.dsn()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping MySQL: %w", err)
	}
	client.db = conn
	return nil
}

func (client *MySQLClient) Close() {
	if client.db != nil {
		client.db.Close()
	}
}

func (client *MySQLClient) DB() *sql.DB {
	return client.db
}
