package producer

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/producer/ports"
)

func Bootstrap() (ports.Producer, error) {
	config := newConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newProducer(config)
}
