package gosqldriver

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	instance *MySQLClient
	once     sync.Once
)

type MySQLClient struct {
	config MySQLClientConfig
	db     *sql.DB
}

func NewMySQLClient(config MySQLClientConfig) (*MySQLClient, error) {
	var err error
	once.Do(func() {
		instance = &MySQLClient{config: config}
		err = instance.connect()
		if err != nil {
			instance = nil
		}
	})
	if instance == nil {
		return nil, fmt.Errorf("failed to initialize MySQLClient: %v", err)
	}
	return instance, nil
}

func (client *MySQLClient) connect() error {
	dsn := client.config.dsn()
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
