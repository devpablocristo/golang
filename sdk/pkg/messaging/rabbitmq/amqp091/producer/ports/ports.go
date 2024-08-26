package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// ProducerService define las operaciones específicas para un productor de RabbitMQ.
type Producer interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	Produce(ctx context.Context, queueName string, replyTo string, corrID string, message any) (string, error)
}

// Config define la configuración específica para un productor de RabbitMQ.
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
