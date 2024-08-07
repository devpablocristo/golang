package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

// StartConsumer inicia el consumidor que recibe mensajes de RabbitMQ
func StartConsumer(channel *amqp091.Channel, queueName string) {
	// Declara la cola
	_, err := channel.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	messages, err := channel.Consume(
		queueName, // nombre de la cola
		"",        // consumer tag
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}

	log.Println("Esperando mensajes. Para salir presiona CTRL+C")
	for msg := range messages {
		log.Printf(" [x] Recibido %s", string(msg.Body))
	}
}
