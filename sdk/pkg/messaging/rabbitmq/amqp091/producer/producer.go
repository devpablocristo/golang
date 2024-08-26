package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/producer/ports"
	"github.com/rabbitmq/amqp091-go"
)

var (
	producerInstance  ports.Producer
	producerOnce      sync.Once
	producerInitError error
)

type rabbitMqProducer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

// newProducer creates a new RabbitMQ producer instance
func newProducer(config ports.Config) (ports.Producer, error) {
	producerOnce.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			producerInitError = fmt.Errorf("failed to connect to RabbitMQ: %v", err)
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			producerInitError = fmt.Errorf("failed to open a channel: %v", err)
			return
		}

		producerInstance = &rabbitMqProducer{connection: conn, channel: ch}
	})

	return producerInstance, producerInitError
}

func GetInstance() (ports.Producer, error) {
	if producerInstance == nil {
		return nil, fmt.Errorf("rabbitmq producer is not initialized")
	}
	return producerInstance, nil
}

func (p *rabbitMqProducer) Produce(ctx context.Context, queueName, replyTo, corrID string, message any) (string, error) {
	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	if corrID == "" {
		corrID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	err = p.channel.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
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

func (p *rabbitMqProducer) Channel() (*amqp091.Channel, error) {
	if p.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return p.channel, nil
}

func (p *rabbitMqProducer) Close() error {
	if err := p.channel.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ channel: %w", err)
	}
	if err := p.connection.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ connection: %w", err)
	}
	return nil
}
