package pkgports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMqClient interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	SendAndReceive(context.Context, string, []byte) ([]byte, error)
}

type RabbitMqConfig interface {
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