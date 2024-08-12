package auth

import (
	pkgports "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/ports"

	ports "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/ports"
)

type messageBroker struct {
	rabbitClient pkgports.RabbitMqClient
}

func NewMessageBroker(client pkgports.RabbitMqClient) ports.MessageBroker {
	return &messageBroker{
		rabbitClient: client,
	}
}

func (b *messageBroker) Login() {}
