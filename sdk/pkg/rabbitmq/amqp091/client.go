package amsgqp

import (
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

var (
	instance RabbitMQClientPort
	once     sync.Once
	errInit  error
)

// RabbitMQClientPort es una interfaz que representa las operaciones que puedes realizar con RabbitMQ.
type RabbitMQClientPort interface {
	Channel() (*amqp091.Channel, error)
	Close() error
}

// rabbitMQClient es una implementación del cliente de RabbitMQ.
type rabbitMQClient struct {
	connection *amqp091.Connection
}

// InitializeRabbitMQClient inicializa una nueva conexión de cliente RabbitMQ.
func InitializeRabbitMQClient(config RabbitMQConfig) error {
	once.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.User, config.Password, config.Host, config.Port, config.VHost)

		conn, err := amqp091.Dial(connString)
		if err != nil {
			errInit = fmt.Errorf("failed to connect to RabbitMQ: %v", err)
			return
		}

		instance = &rabbitMQClient{connection: conn}
	})
	return errInit
}

// GetRabbitMQInstance devuelve la instancia del cliente RabbitMQ.
func GetRabbitMQInstance() (RabbitMQClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("rabbitmq client is not initialized")
	}
	return instance, nil
}

// Channel devuelve un nuevo canal de comunicación con RabbitMQ.
func (client *rabbitMQClient) Channel() (*amqp091.Channel, error) {
	if client.connection == nil {
		return nil, fmt.Errorf("rabbitmq connection is not open")
	}
	return client.connection.Channel()
}

// Close cierra la conexión con RabbitMQ.
func (client *rabbitMQClient) Close() error {
	return client.connection.Close()
}
