package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/pkgports"
)

type rabbitMqBroker struct {
	rabbitClient pkgports.RabbitMqClient
}

func NewRabbitMqBroker(client pkgports.RabbitMqClient) gtwports.MessageBroker {
	return &rabbitMqBroker{
		rabbitClient: client,
	}
}

func (b *rabbitMqBroker) GetUserUUID(ctx context.Context, au *entities.AuthUser) (string, error) {
	// Convertimos AuthUser a LoginRequest
	lr := dto.DomainToLoginResponse(au)

	// Serializar LoginRequest a JSON
	body, err := json.Marshal(lr)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login request: %w", err)
	}

	// Utilizar el cliente de RabbitMQ desde pkg para enviar y recibir mensajes
	responseBody, err := b.rabbitClient.SendAndReceive(ctx, "user_uuid_queue", body)
	if err != nil {
		return "", fmt.Errorf("failed to communicate with RabbitMQ: %w", err)
	}

	// Aquí se asume que la respuesta es un UUID en formato string
	var uuid string
	if err := json.Unmarshal(responseBody, &uuid); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return uuid, nil
}
