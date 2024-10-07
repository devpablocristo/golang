package sdkrabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	"github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/prod-cons/ports"
)

var (
	instance  ports.Messaging
	once      sync.Once
	initError error
)

// service implementa la interfaz ports.Messaging para RabbitMQ.
type service struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	config     ports.Config
}

// newMessaging crea una nueva instancia de RabbitMQ que actúa como productor y consumidor.
func newMessaging(config ports.Config) (ports.Messaging, error) {
	once.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			initError = fmt.Errorf("failed to connect to RabbitMQ: %w", err)
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			initError = fmt.Errorf("failed to open a channel: %w", err)
			conn.Close()
			return
		}

		instance = &service{
			connection: conn,
			channel:    ch,
			config:     config,
		}
	})

	return instance, initError
}

// GetInstance devuelve la instancia única de RabbitMQ como productor y consumidor.
func GetInstance() (ports.Messaging, error) {
	if instance == nil {
		return nil, fmt.Errorf("rabbitmq client instance is not initialized")
	}
	return instance, nil
}

// Publish envía un mensaje al intercambio especificado o directamente a una cola.
func (client *service) Publish(targetType, targetName, routingKey string, body []byte) error {
	var err error
	publishing := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}

	switch targetType {
	case "exchange":
		err = client.channel.Publish(
			targetName, // Exchange
			routingKey, // Routing key
			false,      // Mandatory
			false,      // Immediate
			publishing,
		)
	case "queue":
		err = client.channel.Publish(
			"",         // No exchange (direct to queue)
			targetName, // Queue name
			false,      // Mandatory
			false,      // Immediate
			publishing,
		)
	default:
		return fmt.Errorf("invalid target type: %s", targetType)
	}

	if err != nil {
		return fmt.Errorf("failed to publish message to %s: %w", targetType, err)
	}

	return nil
}

// Subscribe procesa los mensajes de un intercambio específico o una cola específica.
func (client *service) Subscribe(ctx context.Context, targetType, targetName, exchangeType, routingKey string) (<-chan amqp091.Delivery, error) {
	if targetType == "exchange" {
		if err := client.channel.ExchangeDeclare(
			targetName,   // Nombre del intercambio
			exchangeType, // Tipo de intercambio (direct, topic, fanout, etc.)
			true,         // Durable
			false,        // Auto-deleted
			false,        // Internal
			false,        // No-wait
			nil,          // Arguments
		); err != nil {
			return nil, fmt.Errorf("failed to declare exchange: %w", err)
		}
	}

	queue, err := client.channel.QueueDeclare(
		targetName, // Nombre de la cola
		true,       // Durable
		false,      // Delete when unused
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	if targetType == "exchange" {
		if err := client.channel.QueueBind(
			queue.Name, // Nombre de la cola
			routingKey, // Clave de enrutamiento
			targetName, // Nombre del intercambio
			false,      // No-wait
			nil,        // Arguments
		); err != nil {
			return nil, fmt.Errorf("failed to bind queue: %w", err)
		}
	}

	msgs, err := client.channel.Consume(
		queue.Name,                 // Nombre de la cola
		"",                         // Consumer
		client.config.GetAutoAck(), // Auto-acknowledge
		client.config.GetExclusive(),
		client.config.GetNoLocal(),
		client.config.GetNoWait(),
		nil, // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	// Crear un canal para filtrar mensajes con cancelación de contexto
	filteredMsgs := make(chan amqp091.Delivery)

	go func() {
		defer close(filteredMsgs)
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgs:
				if !ok {
					return
				}
				filteredMsgs <- msg
			}
		}
	}()

	return filteredMsgs, nil
}

// SetupExchangeAndQueue configura el intercambio y la cola en RabbitMQ.
func (client *service) SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error {
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

// Close cierra de manera segura la conexión de RabbitMQ.
func (client *service) Close() error {
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
