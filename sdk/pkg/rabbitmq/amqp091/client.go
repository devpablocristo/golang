package amsgqp

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/pkgports"
	"github.com/rabbitmq/amqp091-go"
)

var (
	instance pkgports.RabbitMqClient
	once     sync.Once
	errInit  error
)

type rabbitMqClient struct {
	connection *amqp091.Connection
}

// InitializeRabbitMQClient inicializa una conexión única a RabbitMQ.
func InitializeRabbitMQClient(config pkgports.RabbitMqConfig) error {
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
func GetRabbitMQInstance() (pkgports.RabbitMqClient, error) {
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

// SendAndReceive envía un mensaje a RabbitMQ y espera una respuesta en la misma cola.
func (client *rabbitMqClient) SendAndReceive(ctx context.Context, queueName string, message []byte) ([]byte, error) {
	// Abre un canal
	ch, err := client.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	// Declara la cola donde se enviará el mensaje
	q, err := ch.QueueDeclare(
		queueName, // Nombre de la cola
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare RabbitMQ queue: %w", err)
	}

	// Genera un correl_id para rastrear la respuesta
	corrId := fmt.Sprintf("%d", time.Now().UnixNano())

	// Declara una cola temporal para recibir la respuesta
	replyQueue, err := ch.QueueDeclare(
		"",    // Nombre de la cola (vacío para que sea temporal)
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare RabbitMQ reply queue: %w", err)
	}

	// Publica el mensaje en la cola original
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			Body:          message,
			ReplyTo:       replyQueue.Name,
			CorrelationId: corrId,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to publish message to RabbitMQ: %w", err)
	}

	// Consume el mensaje de respuesta
	msgs, err := ch.Consume(
		replyQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	// Esperar y recibir la respuesta con el mismo correl_id
	for d := range msgs {
		if d.CorrelationId == corrId {
			return d.Body, nil
		}
	}

	return nil, fmt.Errorf("no response received for correlation ID: %s", corrId)
}
