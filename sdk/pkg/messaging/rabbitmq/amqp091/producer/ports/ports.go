package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Producer define las operaciones específicas para un productor de RabbitMQ.
type Producer interface {
	// Channel devuelve el canal actual de RabbitMQ.
	Channel() (*amqp091.Channel, error)

	// Close cierra de manera segura el productor de RabbitMQ.
	Close() error

	// Produce envía un mensaje a la cola especificada con una opción de reply-to y ID de correlación.
	Produce(ctx context.Context, queueName string, replyTo string, corrID string, message any) (string, error)

	// ProduceWithRetry envía un mensaje con reintentos en caso de fallo.
	ProduceWithRetry(ctx context.Context, queueName string, replyTo string, corrID string, message any, maxRetries int) (string, error)
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

	GetExchange() string
	SetExchange(exchange string)

	GetExchangeType() string
	SetExchangeType(exchangeType string)

	IsDurable() bool
	SetDurable(durable bool)

	IsAutoDelete() bool
	SetAutoDelete(autoDelete bool)

	IsInternal() bool
	SetInternal(internal bool)

	IsNoWait() bool
	SetNoWait(noWait bool)

	Validate() error
}
