package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// ConsumerService define las operaciones específicas para un consumidor de RabbitMQ.
type Consumer interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	Consume(ctx context.Context, queueName string, corrID string) ([]byte, string, error)
}

// ConsumerConfig define la configuración específica para un consumidor de RabbitMQ.
type Config interface {
	GetHost() string
	SetHost(host string)
	GetPort() int
	SetPort(port int)
	GetUser() string
	SetUser(user string)
	GetPassword() string
	SetPassword(password string)
	GetVHost() string
	SetVHost(vhost string)
	Validate() error
}
