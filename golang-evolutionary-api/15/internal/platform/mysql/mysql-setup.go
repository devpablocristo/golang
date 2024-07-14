package mysqlsetup

import (
	gosqldriver "api/pkg/mysql/go-sql-driver"
)

func NewMySQLSetup() (*gosqldriver.MySQLClient, error) {
	config := gosqldriver.MySQLClientConfig{
		User:     "user",
		Password: "password",
		Host:     "mysql",
		Port:     "3306",
		Database: "inventory",
	}
	return gosqldriver.NewMySQLClient(config)
}
