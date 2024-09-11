package producer

import (
	"fmt"

	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/producer/ports"
)

// Bootstrap inicializa una nueva instancia de Producer con configuración de Viper.
func Bootstrap() (ports.Producer, error) {
	// Recuperar la configuración desde Viper
	host := viper.GetString("RABBITMQ_HOST")
	port := viper.GetInt("RABBITMQ_PORT")
	user := viper.GetString("RABBITMQ_USER")
	password := viper.GetString("RABBITMQ_PASSWORD")
	vhost := viper.GetString("RABBITMQ_VHOST")
	exchange := viper.GetString("RABBITMQ_EXCHANGE")
	exchangeType := viper.GetString("RABBITMQ_EXCHANGE_TYPE")

	// Configuraciones booleanas con valores predeterminados
	durable := viper.GetBool("RABBITMQ_DURABLE")
	autoDelete := viper.GetBool("RABBITMQ_AUTO_DELETE")
	internal := viper.GetBool("RABBITMQ_INTERNAL")
	noWait := viper.GetBool("RABBITMQ_NO_WAIT")

	// Crear una nueva configuración
	config := newConfig(host, port, user, password, vhost, exchange, exchangeType, durable, autoDelete, internal, noWait)

	// Validar la configuración
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	// Crear una nueva instancia de productor
	producer, err := newProducer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create new producer: %w", err)
	}

	return producer, nil
}
