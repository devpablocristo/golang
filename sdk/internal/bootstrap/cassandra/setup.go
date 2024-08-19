package cassandrasetup

import (
	"github.com/spf13/viper"

	csdgocsl "github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql/portspkg"
)

func NewCassandraInstance() (portspkg.CassandraClient, error) {
	// Crear una instancia de configuración de Cassandra usando NewCassandraConfig
	config := csdgocsl.NewCassandraConfig(
		viper.GetStringSlice("CASSANDRA_HOSTS"),
		viper.GetString("CASSANDRA_KEYSPACE"),
		viper.GetString("CASSANDRA_USERNAME"),
		viper.GetString("CASSANDRA_PASSWORD"),
	)

	// Validar la configuración antes de inicializar el cliente
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Inicializar el cliente de Cassandra con la configuración
	if err := csdgocsl.InitializeCassandraClient(config); err != nil {
		return nil, err
	}

	// Obtener la instancia de Cassandra
	return csdgocsl.GetCassandraInstance()
}
