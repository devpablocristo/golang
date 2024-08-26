package auth

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/internal/core/auth/entities"

	ports "github/devpablocristo/golang/sdk/internal/core/auth/ports"
)

type RabbitMqBroker struct {
	ucs ports.AuthUseCases
}

func NewRabbitMqBroker(u ports.AuthUseCases) *RabbitMqBroker {
	return &RabbitMqBroker{
		ucs: u,
	}
}

func (b *RabbitMqBroker) GetUserUUID(ctx context.Context, lc *entities.LoginCredentials) (string, error) {
	// Preparar el mensaje de solicitud
	// loginCredentials := dto.DomainToLoginResponse(lc)

	// // Definir las colas
	// queueName := "getuseruuid_req_queue"
	// replyTo := "getuseruuid_res_queue"

	// // Enviar el mensaje usando el Producer del SDK
	// corrID, err := b.broker.Produce(ctx, queueName, replyTo, "", loginCredentials)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to send login request: %w", err)
	// }

	// // Consumir la respuesta de la cola de respuestas
	// responseBody, returnedCorrID, err := b.broker.Consume(ctx, replyTo, corrID)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to consume response from RabbitMQ: %w", err)
	// }

	// // Verificar si el `corrID` recibido coincide con el que se envió
	// if corrID != returnedCorrID {
	// 	return "", fmt.Errorf("mismatched correlation ID: expected %s but got %s", corrID, returnedCorrID)
	// }

	// // Procesar la respuesta (por ejemplo, deserializar UUID)
	// var uuid string
	// if err := json.Unmarshal(responseBody, &uuid); err != nil {
	// 	return "", fmt.Errorf("failed to unmarshal response: %w", err)
	// }

	uuid := "12345"
	return uuid, nil
}
