package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Messaging define las operaciones para un sistema de mensajería RabbitMQ.
type Messaging interface {
	Publish(targetType, targetName, routingKey string, body []byte) error
	Subscribe(ctx context.Context, targetType, targetName, exchangeType, routingKey string) (<-chan amqp091.Delivery, error)
	SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error
	Close() error
}

// Config define la configuración específica para RabbitMQ.
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
