package amqpsetup

import (
	"github.com/spf13/viper"

	amsgqp "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091"
	pkgports "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/ports"
)

func NewRabbitMQInstance() (pkgports.RabbitMqClient, error) {
	config := amsgqp.NewRabbitMqConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := amsgqp.InitializeRabbitMQClient(config); err != nil {
		return nil, err
	}

	return amsgqp.GetRabbitMQInstance()
}
