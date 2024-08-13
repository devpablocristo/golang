package rabbitpkg

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/portspkg"
	"github.com/rabbitmq/amqp091-go"
)

var (
	instance portspkg.RabbitMqClient
	once     sync.Once
	errInit  error
)

type rabbitMqClient struct {
	connection *amqp091.Connection
}

// InitializeRabbitMQClient inicializa una conexión única a RabbitMQ.
func InitializeRabbitMQClient(config portspkg.RabbitMqConfig) error {
	once.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			errInit = fmt.Errorf("failed to connect to RabbitMQ: %v", err)
			return
		}

		instance = &rabbitMqClient{connection: conn}
	})
	return errInit
}

// GetRabbitMQInstance devuelve la instancia del cliente RabbitMQ.
func GetRabbitMQInstance() (portspkg.RabbitMqClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("rabbitmq client is not initialized")
	}
	return instance, nil
}

// Channel devuelve un nuevo canal de comunicación con RabbitMQ.
func (client *rabbitMqClient) Channel() (*amqp091.Channel, error) {
	if client.connection == nil {
		return nil, fmt.Errorf("rabbitmq connection is not open")
	}
	return client.connection.Channel()
}

// Close cierra la conexión con RabbitMQ.
func (client *rabbitMqClient) Close() error {
	return client.connection.Close()
}

func (client *rabbitMqClient) Produce(ctx context.Context, queueName string, replyTo string, corrID string, message interface{}) (string, error) {
	// Abre un canal
	ch, err := client.Channel()
	if err != nil {
		return "", fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	// Convertir el mensaje a formato JSON
	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	if corrID == "" { // Genera un correl_id para rastrear la respuesta
		corrID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Publica el mensaje en la cola original
	err = ch.PublishWithContext(ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: corrID,
			ReplyTo:       replyTo,
		})
	if err != nil {
		return "", fmt.Errorf("failed to publish message to RabbitMQ: %w", err)
	}

	return corrID, nil
}

func (client *rabbitMqClient) Consume(ctx context.Context, queueName string, corrID string) ([]byte, string, error) {
	ch, err := client.Channel()
	if err != nil {
		return nil, "", fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	for msg := range msgs {
		if corrID == "" {
			corrID = msg.CorrelationId
		}

		if msg.CorrelationId == corrID {
			return msg.Body, corrID, nil
		}
	}

	return nil, "", fmt.Errorf("no response received for correlation ID: %s", corrID)
}
