package sdkrabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	"github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/consumer/ports"
)

var (
	consumerInstance  ports.Consumer
	consumerOnce      sync.Once
	consumerInitError error
)

// rabbitMqConsumer implementa la interfaz ports.Consumer para RabbitMQ.
type rabbitMqConsumer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	config     ports.Config
}

// newConsumer crea una nueva instancia de consumidor de RabbitMQ utilizando el patrón Singleton.
func newConsumer(config ports.Config) (ports.Consumer, error) {
	consumerOnce.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			consumerInitError = fmt.Errorf("failed to connect to RabbitMQ: %w", err)
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			consumerInitError = fmt.Errorf("failed to open a channel: %w", err)
			conn.Close()
			return
		}

		consumerInstance = &rabbitMqConsumer{
			connection: conn,
			channel:    ch,
			config:     config,
		}
	})

	return consumerInstance, consumerInitError
}

// GetConsumerInstance devuelve la instancia única del consumidor de RabbitMQ.
func GetConsumerInstance() (ports.Consumer, error) {
	if consumerInstance == nil {
		return nil, fmt.Errorf("rabbitmq consumer is not initialized")
	}
	return consumerInstance, nil
}

// Consume procesa los mensajes de la cola especificada.
// Retorna el primer mensaje que coincide con el corrID proporcionado.
// Si corrID está vacío, retorna el primer mensaje recibido.
func (client *rabbitMqConsumer) Consume(ctx context.Context, queueName, corrID string) ([]byte, string, error) {
	msgs, err := client.channel.Consume(
		queueName, "", client.config.GetAutoAck(), client.config.GetExclusive(),
		client.config.GetNoLocal(), client.config.GetNoWait(), nil,
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil, "", ctx.Err()
		case msg, ok := <-msgs:
			if !ok {
				return nil, "", fmt.Errorf("message channel closed")
			}
			if corrID == "" || msg.CorrelationId == corrID {
				return msg.Body, msg.CorrelationId, nil
			}
		}
	}
}

// SetupExchangeAndQueue configura el intercambio y la cola en RabbitMQ.
func (client *rabbitMqConsumer) SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error {
	if err := client.channel.ExchangeDeclare(
		exchangeName, // Nombre del intercambio
		exchangeType, // Tipo de intercambio (direct, topic, fanout, etc.)
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	if _, err := client.channel.QueueDeclare(
		queueName, // Nombre de la cola
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	); err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	if err := client.channel.QueueBind(
		queueName,    // Nombre de la cola
		routingKey,   // Clave de enrutamiento
		exchangeName, // Nombre del intercambio
		false,        // No-wait
		nil,          // Arguments
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}

// Close cierra de manera segura el consumidor.
func (client *rabbitMqConsumer) Close() error {
	var errs []error

	if err := client.channel.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ channel: %w", err))
	}

	if err := client.connection.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ connection: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing consumer: %v", errs)
	}

	return nil
}

// Channel devuelve el canal actual de RabbitMQ.
func (client *rabbitMqConsumer) Channel() (*amqp091.Channel, error) {
	if client.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return client.channel, nil
}
