package amqp

import (
	"fmt"
	"os"
)

// Config representa la configuración para conectar con RabbitMQ.
type Config struct {
	URI       string
	QueueName string
}

// LoadConfig carga la configuración desde variables de entorno.
func LoadConfig() (*Config, error) {
	uri := os.Getenv("RABBITMQ_URI")
	queueName := os.Getenv("RABBITMQ_QUEUE")

	if uri == "" || queueName == "" {
		return nil, fmt.Errorf("rabbitmq configuration: URI and QueueName must be set")
	}

	return &Config{
		URI:       uri,
		QueueName: queueName,
	}, nil
}