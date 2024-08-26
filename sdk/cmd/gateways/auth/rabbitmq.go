package auth

import (
	"context"
	"encoding/json"
	"fmt"

	dto "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	entities "github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/auth/ports"
)

type RabbitMq struct {
	ucs    ports.AuthUseCases
	broker ports.Service // Asumimos que ports.Service es la interfaz del cliente RabbitMQ
}

func NewRabbitMq(u ports.AuthUseCases, broker ports.Service) *RabbitMq {
	return &RabbitMq{
		ucs:    u,
		broker: broker,
	}
}

func (b *RabbitMq) GetUserUUID(ctx context.Context, lc *entities.LoginCredentials) (string, error) {
	// Preparar el mensaje de solicitud
	loginCredentials := dto.DomainToLoginResponse(lc)

	// Definir las colas
	queueName := "getuseruuid_req_queue"
	replyTo := "getuseruuid_res_queue"

	// Enviar el mensaje usando el Producer del SDK
	corrID, err := b.broker.Produce(ctx, queueName, replyTo, "", loginCredentials)
	if err != nil {
		return "", fmt.Errorf("failed to send login request: %w", err)
	}

	// Consumir la respuesta de la cola de respuestas
	responseBody, returnedCorrID, err := b.broker.Consume(ctx, replyTo, corrID)
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
