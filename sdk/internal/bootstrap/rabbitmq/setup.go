package amqpsetup

import (
	"github.com/spf13/viper"

	rabbitpkg "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091"
	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/portspkg"
)

func NewRabbitMQInstance() (portspkg.RabbitMqClient, error) {
	config := rabbitpkg.NewRabbitMqConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := rabbitpkg.InitializeRabbitMQClient(config); err != nil {
		return nil, err
	}

	return rabbitpkg.GetRabbitMQInstance()
}
