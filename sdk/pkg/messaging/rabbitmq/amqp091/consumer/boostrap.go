package consumer

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/consumer/ports"
)

func Bootstrap() (ports.Consumer, error) {
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

	return newConsumer(config)
}
