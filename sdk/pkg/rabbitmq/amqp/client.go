package amqp

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

// RabbitMQPublisherPort define la interfaz para publicar mensajes.
type RabbitMQPublisherPort interface {
	Publish(body []byte) error
	Close()
}

// RabbitMQConsumerPort define la interfaz para consumir mensajes.
type RabbitMQConsumerPort interface {
	Consume(handler func(amqp.Delivery)) error
	Close()
}

var (
	publisherInstance RabbitMQPublisherPort
	consumerInstance  RabbitMQConsumerPort
	once              sync.Once
	initErr           error
)

// RabbitMQClient gestiona las conexiones a RabbitMQ.
type RabbitMQClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

// InitializeRabbitMQ inicializa el cliente RabbitMQ.
func InitializeRabbitMQ(config *Config) error {
	once.Do(func() {
		client := &RabbitMQClient{queueName: config.QueueName}
		initErr = client.connect(config.URI)
		if initErr != nil {
			publisherInstance = nil
			consumerInstance = nil
		} else {
			publisherInstance = client
			consumerInstance = client
		}
	})
	return initErr
}

// GetPublisherInstance devuelve la instancia del publicador.
func GetPublisherInstance() (RabbitMQPublisherPort, error) {
	if publisherInstance == nil {
		return nil, fmt.Errorf("RabbitMQPublisher: client is not initialized")
	}
	return publisherInstance, nil
}

// GetConsumerInstance devuelve la instancia del consumidor.
func GetConsumerInstance() (RabbitMQConsumerPort, error) {
	if consumerInstance == nil {
		return nil, fmt.Errorf("RabbitMQConsumer: client is not initialized")
	}
	return consumerInstance, nil
}

// connect establece la conexión con RabbitMQ.
func (client *RabbitMQClient) connect(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	client.connection = conn
	client.channel = ch
	log.Println("RabbitMQ connected successfully")
	return nil
}

// Publish envía un mensaje al broker de RabbitMQ.
func (client *RabbitMQClient) Publish(body []byte) error {
	queue, err := client.channel.QueueDeclare(
		client.queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = client.channel.Publish(
		"",         // Exchange
		queue.Name, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	log.Printf("Message published to queue %s", client.queueName)
	return nil
}

// Consume inicia el consumo de mensajes desde una cola específica en RabbitMQ.
func (client *RabbitMQClient) Consume(handler func(amqp.Delivery)) error {
	msgs, err := client.channel.Consume(
		client.queueName,
		"",    // Consumer
		true,  // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			handler(msg)
			log.Printf("Message consumed from queue %s", client.queueName)
		}
	}()

	return nil
}

// Close cierra la conexión y el canal de RabbitMQ.
func (client *RabbitMQClient) Close() {
	err := client.channel.Close()
	if err != nil {
		log.Printf("Error closing channel: %v", err)
	}
	err = client.connection.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}