package postgresqlsetup

import (
	"github.com/spf13/viper"
)

func NewPostgreSQLInstance() (pqpostgresql.PostgreSQLClientPort, error) {
	config := pqpostgresql.PostgreSQLClientConfig{
		User:     viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		DBName:   viper.GetString("POSTGRES_DB"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := pqpostgresql.InitializePostgreSQLClient(config); err != nil {
		return nil, err
	}

	return pqpostgresql.GetPostgreSQLInstance()
}
