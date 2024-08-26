package pkgrabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type rabbitMqClient struct {
	connection *amqp091.Connection
}

func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			initError = fmt.Errorf("failed to connect to RabbitMQ: %v", err)
			return
		}

		instance = &rabbitMqClient{connection: conn}
	})
	return instance, initError
}

func GetInstance() (ports.Service, error) {
	if instance == nil {
		return nil, fmt.Errorf("rabbitmq client is not initialized")
	}
	return instance, nil
}

func (client *rabbitMqClient) Channel() (*amqp091.Channel, error) {
	if client.connection == nil {
		return nil, fmt.Errorf("rabbitmq connection is not open")
	}
	return client.connection.Channel()
}

func (client *rabbitMqClient) Close() error {
	if client.connection != nil {
		return client.connection.Close()
	}
	return nil
}

func (client *rabbitMqClient) Produce(ctx context.Context, queueName, replyTo, corrID string, message any) (string, error) {
	ch, err := client.Channel()
	if err != nil {
		return "", fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	if corrID == "" {
		corrID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

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

func (client *rabbitMqClient) Consume(ctx context.Context, queueName, corrID string) ([]byte, string, error) {
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
