package mysqlsetup

import (
	"github.com/spf13/viper"

	gosqldriver "github.com/devpablocristo/golang-sdk/pkg/mysql/go-sql-driver"
)

func NewMySQLInstance() (gosqldriver.MySQLClientPort, error) {
	config := gosqldriver.MySQLClientConfig{
		User:     viper.GetString("MYSQL_USER"),
		Password: viper.GetString("MYSQL_PASSWORD"),
		Host:     viper.GetString("MYSQL_HOST"),
		Port:     viper.GetString("MYSQL_PORT"),
		Database: viper.GetString("MYSQL_DATABASE"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := gosqldriver.InitializeMySQLClient(config); err != nil {
		return nil, err
	}

	return gosqldriver.GetMySQLInstance()
}
