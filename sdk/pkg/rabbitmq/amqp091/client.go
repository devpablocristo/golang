package amsgqp

import (
	"fmt"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/rabbitmq/amqp091/port"
	"github.com/rabbitmq/amqp091-go"
)

var (
	instance port.RabbitMqClient
	once     sync.Once
	errInit  error
)

type rabbitMqClient struct {
	connection *amqp091.Connection
}

func InitializeRabbitMQClient(config port.RabbitMqConfig) error {
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
func GetRabbitMQInstance() (port.RabbitMqClient, error) {
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
