package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	exchange   string
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

		// Set up exchange (configurable)
		err = ch.ExchangeDeclare(
			config.GetExchange(),     // Exchange name
			config.GetExchangeType(), // Exchange type (direct, topic, fanout, etc.)
			config.IsDurable(),       // Durable
			config.IsAutoDelete(),    // Auto-deleted
			config.IsInternal(),      // Internal
			config.IsNoWait(),        // No-wait
			nil,                      // Arguments
		)
		if err != nil {
			producerInitError = fmt.Errorf("failed to declare exchange: %v", err)
			return
		}

		// Set up publisher confirms
		err = ch.Confirm(false)
		if err != nil {
			producerInitError = fmt.Errorf("failed to put channel into confirm mode: %v", err)
			return
		}

		producerInstance = &rabbitMqProducer{connection: conn, channel: ch, exchange: config.GetExchange()}
	})

	return producerInstance, producerInitError
}

// GetInstance returns the singleton instance of the RabbitMQ producer
func GetInstance() (ports.Producer, error) {
	if producerInstance == nil {
		return nil, fmt.Errorf("rabbitmq producer is not initialized")
	}
	return producerInstance, nil
}

// Produce sends a message to the specified queue with optional reply-to and correlation ID
func (p *rabbitMqProducer) Produce(ctx context.Context, queueName, replyTo, corrID string, message any) (string, error) {
	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	if corrID == "" {
		corrID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	err = p.channel.PublishWithContext(ctx,
		p.exchange, // Exchange name
		queueName,  // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: corrID,
			ReplyTo:       replyTo,
		})
	if err != nil {
		return "", fmt.Errorf("failed to publish message to RabbitMQ: %w", err)
	}

	// Confirm message was received by RabbitMQ
	ack, nack := make(chan uint64), make(chan uint64)
	p.channel.NotifyConfirm(ack, nack)

	select {
	case <-ack:
		log.Println("Message acknowledged by RabbitMQ")
	case <-nack:
		return "", fmt.Errorf("message not acknowledged by RabbitMQ")
	case <-ctx.Done():
		return "", ctx.Err()
	}

	return corrID, nil
}

// ProduceWithRetry sends a message with retries in case of failure
func (p *rabbitMqProducer) ProduceWithRetry(ctx context.Context, queueName, replyTo, corrID string, message any, maxRetries int) (string, error) {
	var err error
	for i := 0; i < maxRetries; i++ {
		corrID, err = p.Produce(ctx, queueName, replyTo, corrID, message)
		if err == nil {
			return corrID, nil
		}
		log.Printf("Retry %d/%d failed: %v", i+1, maxRetries, err)
	}
	return "", fmt.Errorf("max retries reached: %w", err)
}

// Channel returns the current RabbitMQ channel
func (p *rabbitMqProducer) Channel() (*amqp091.Channel, error) {
	if p.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return p.channel, nil
}

// Close gracefully closes the RabbitMQ producer
func (p *rabbitMqProducer) Close() error {
	if err := p.channel.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ channel: %w", err)
	}
	if err := p.connection.Close(); err != nil {
		return fmt.Errorf("failed to close RabbitMQ connection: %w", err)
	}
	return nil
}
