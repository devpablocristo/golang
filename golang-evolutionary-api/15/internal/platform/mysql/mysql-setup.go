package mysqlsetup

import (
	gosqldriver "api/pkg/mysql/go-sql-driver"
)

func NewMySQLSetup() (*gosqldriver.MySQLClient, error) {
	config := gosqldriver.MySQLClientConfig{
		User:     "DB_USER",
		Password: "DB_PASSWORD",
		Host:     "DB_HOST",
		Port:     "DB_PORT",
		Database: "DB_NAME",
	}
	return gosqldriver.NewMySQLClient(config)
}
