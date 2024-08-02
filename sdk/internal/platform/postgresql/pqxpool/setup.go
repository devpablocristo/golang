package postgresqlsetup

import (
	"github.com/spf13/viper"

	pgxsdk "github.com/devpablocristo/golang-sdk/pkg/postgresql/pgxpool"
)

func NewPostgreSQLInstance() (pgxsdk.PostgreSQLClientPort, error) {
	config := pgxsdk.PostgreSQLClientConfig{
		User:     viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		DBName:   viper.GetString("POSTGRES_NAME"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := pgxsdk.InitializePostgreSQLClient(config); err != nil {
		return nil, err
	}

	return pgxsdk.GetPostgreSQLInstance()
}
