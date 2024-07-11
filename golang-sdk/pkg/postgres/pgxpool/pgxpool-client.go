package pgxpostgres

import (
	"fmt"
	"sync"

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
