package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/portspkg"
)

type rabbitMqBroker struct {
	client portspkg.RabbitMqClient
}

func NewRabbitMqBroker(client portspkg.RabbitMqClient) gtwports.MessageBroker {
	return &rabbitMqBroker{
		client: client,
	}
}

func (b *rabbitMqBroker) GetUserUUID(ctx context.Context, lc *entities.LogingCredentials) (string, error) {
	lr := dto.DomainToLoginResponse(lc)

	body, err := json.Marshal(lr)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login request: %w", err)
	}

	corrId, err := b.client.Produce(ctx, "user_uuid_queue", body)
	if err != nil {
		return "", fmt.Errorf("failed to communicate with RabbitMQ: %w", err)
	}

	responseBody, err := b.client.Consume(ctx, "user_uuid_reply_queue", corrId)
	if err != nil {
		return "", fmt.Errorf("failed to consume message from RabbitMQ: %w", err)
	}

	// Deserializar la respuesta que se espera sea solo un UUID en formato string
	var uuid string
	if err := json.Unmarshal(responseBody, &uuid); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Verificar si el UUID está vacío, lo que indicaría que el usuario no existe
	if uuid == "" {
		return "", fmt.Errorf("user does not exist")
	}

	return uuid, nil
}
