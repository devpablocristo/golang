package portspkg

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	Produce(context.Context, string, string, string, any) (string, error)
	Consume(context.Context, string, string) ([]byte, string, error)
}

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
