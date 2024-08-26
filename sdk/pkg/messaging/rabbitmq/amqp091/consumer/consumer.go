package consumer

import (
	"context"
	"fmt"
	"sync"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/consumer/ports"
	"github.com/rabbitmq/amqp091-go"
)

var (
	consumerInstance  ports.Consumer
	consumerOnce      sync.Once
	consumerInitError error
)

type rabbitMqConsumer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

// newConsumer creates a new RabbitMQ consumer instance
func newConsumer(config ports.Config) (ports.Consumer, error) {
	consumerOnce.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			consumerInitError = fmt.Errorf("failed to connect to RabbitMQ: %v", err)
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			consumerInitError = fmt.Errorf("failed to open a channel: %v", err)
			return
		}

		consumerInstance = &rabbitMqConsumer{connection: conn, channel: ch}
	})

	return consumerInstance, consumerInitError
}

func GetConsumerInstance() (ports.Consumer, error) {
	if consumerInstance == nil {
		return nil, fmt.Errorf("rabbitmq consumer is not initialized")
	}
	return consumerInstance, nil
}

func (client *rabbitMqConsumer) Consume(ctx context.Context, queueName, corrID string) ([]byte, string, error) {
	msgs, err := client.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	for msg := range msgs {
		if corrID == "" || msg.CorrelationId == corrID {
			return msg.Body, msg.CorrelationId, nil
		}
	}

	return nil, "", fmt.Errorf("no response received for correlation ID: %s", corrID)
}

func (client *rabbitMqConsumer) Channel() (*amqp091.Channel, error) {
	if client.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return client.channel, nil
}

func (client *rabbitMqConsumer) Close() error {
	if err := client.channel.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ channel: %w", err)
	}
	if err := client.connection.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ connection: %w", err)
	}
	return nil
}
