package user

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/devpablocristo/golang/sdk/gateways/user/portsgtw"
// 	mdw "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
// 	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/portspkg"
// )

// type rabbitMqBroker struct {
// 	client portspkg.Service
// }

// func NewRabbitMqBroker(client portspkg.Service) portsgtw.MessageBroker {
// 	return &rabbitMqBroker{
// 		client: client,
// 	}
// }

// func (b *rabbitMqBroker) SendUser(ctx context.Context) error {
// 	// Consumir mensaje de la cola de solicitudes
// 	messageBody, corrID, err := b.client.Consume(ctx, "getuseruuid_req_queue", "")
// 	if err != nil {
// 		return fmt.Errorf("failed to consume message from RabbitMQ: %w", err)
// 	}

// 	// Deserializar las credenciales del usuario desde el mensaje
// 	var lc mdw.LoginRequest
// 	if err := json.Unmarshal(messageBody, &lc); err != nil {
// 		return fmt.Errorf("failed to unmarshal login credentials: %w", err)
// 	}

// 	// Procesar la solicitud para obtener el UUID del usuario (lógica específica)
// 	uuid := "simulated-uuid" // Aquí iría la lógica real para obtener el UUID basado en lc

// 	// Preparar la respuesta
// 	responseBody, err := json.Marshal(uuid)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal response: %w", err)
// 	}

// 	// Enviar la respuesta a la cola especificada en ReplyTo del mensaje original
// 	_, err = b.client.Produce(ctx, "getuseruuid_res_queue", "", corrID, responseBody)
// 	if err != nil {
// 		return fmt.Errorf("failed to send response to RabbitMQ: %w", err)
// 	}

// 	return nil
// }