package amqpsetup

import (
	"github.com/spf13/viper"

	amsgqp "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091"
)

func NewRabbitMQInstance() (amsgqp.RabbitMQClientPort, error) {
	config := amsgqp.RabbitMQConfig{
		Host:     viper.GetString("RABBITMQ_HOST"),
		Port:     viper.GetInt("RABBITMQ_PORT"),
		User:     viper.GetString("RABBITMQ_USER"),
		Password: viper.GetString("RABBITMQ_PASSWORD"),
		VHost:    viper.GetString("RABBITMQ_VHOST"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := amsgqp.InitializeRabbitMQClient(config); err != nil {
		return nil, err
	}

	return amsgqp.GetRabbitMQInstance()
}
