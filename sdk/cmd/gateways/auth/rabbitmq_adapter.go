package auth

import (
	pkgport "github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/port"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/port"
)

type messageBroker struct {
	rabbitClient pkgport.RabbitMqClient
}

// NewMessageBroker crea un nuevo messageBroker
func NewMessageBroker(client pkgport.RabbitMqClient) port.MessageBroker {
	return &messageBroker{
		rabbitClient: client,
	}
}

func (b *messageBroker) Login() {}
