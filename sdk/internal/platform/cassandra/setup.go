package cassandrasetup

import (
	"github.com/spf13/viper"

	csdgocsl "github.com/devpablocristo/golang-sdk/pkg/cassandra/gocql"
)

func NewCassandraInstance() (csdgocsl.CassandraClientPort, error) {
	config := csdgocsl.CassandraConfig{
		Hosts:    viper.GetStringSlice("CASSANDRA_HOSTS"),
		Keyspace: viper.GetString("CASSANDRA_KEYSPACE"),
		Username: viper.GetString("CASSANDRA_USERNAME"),
		Password: viper.GetString("CASSANDRA_PASSWORD"),
	}

	if err := csdgocsl.InitializeCassandraClient(config); err != nil {
		return nil, err
	}

	return csdgocsl.GetCassandraInstance()
}
