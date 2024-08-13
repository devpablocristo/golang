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
	rabbitClient portspkg.RabbitMqClient
}

func NewRabbitMqBroker(client portspkg.RabbitMqClient) gtwports.MessageBroker {
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
