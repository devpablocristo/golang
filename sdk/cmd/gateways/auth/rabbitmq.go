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

func (b *rabbitMqBroker) GetUserUUID(ctx context.Context, lc *entities.LoginCredentials) (string, error) {
	// Preparar el mensaje de solicitud
	loginCredentials := dto.DomainToLoginResponse(lc)

	// Definir las colas
	queueName := "getuseruuid_req_queue"
	replyTo := "getuseruuid_res_queue"

	// Enviar el mensaje usando el Producer del SDK
	corrID, err := b.client.Produce(ctx, queueName, replyTo, "", loginCredentials)
	if err != nil {
		return "", fmt.Errorf("failed to send login request: %w", err)
	}

	// Consumir la respuesta de la cola de respuestas
	responseBody, returnedCorrID, err := b.client.Consume(ctx, replyTo, corrID)
	if err != nil {
		return "", fmt.Errorf("failed to consume response from RabbitMQ: %w", err)
	}

	// Verificar si el `corrID` recibido coincide con el que se envió
	if corrID != returnedCorrID {
		return "", fmt.Errorf("mismatched correlation ID: expected %s but got %s", corrID, returnedCorrID)
	}

	// Procesar la respuesta (por ejemplo, deserializar UUID)
	var uuid string
	if err := json.Unmarshal(responseBody, &uuid); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return uuid, nil
}
