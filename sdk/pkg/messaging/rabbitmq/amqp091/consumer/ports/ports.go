package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Consumer define las operaciones específicas para un consumidor de RabbitMQ.
type Consumer interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	Consume(ctx context.Context, queueName, corrID string) ([]byte, string, error)
	SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error
	GetConnection() *amqp091.Connection
}

// Config define la configuración específica para un consumidor de RabbitMQ.
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

	GetQueue() string
	SetQueue(queue string)

	GetAutoAck() bool
	SetAutoAck(autoAck bool)

	GetExclusive() bool
	SetExclusive(exclusive bool)

	GetNoLocal() bool
	SetNoLocal(noLocal bool)

	GetNoWait() bool
	SetNoWait(noWait bool)

	Validate() error
}
