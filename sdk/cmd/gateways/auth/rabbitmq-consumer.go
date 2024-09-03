package auth

import (
	ports "github.com/devpablocristo/golang/sdk/internal/core/auth/ports"
	rabbitports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/consumer/ports"
)

type RabbitMq struct {
	ucs    ports.UseCases
	broker rabbitports.Consumer
}

func NewRabbitMq(u ports.UseCases, broker rabbitports.Consumer) *RabbitMq {
	return &RabbitMq{
		ucs:    u,
		broker: broker,
	}
}
