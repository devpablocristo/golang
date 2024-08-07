package messaging

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

// StartProducer inicia el productor que envía mensajes a RabbitMQ
func StartProducer(channel *amqp091.Channel, queueName string) {
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

	// Publicar mensajes
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Mensaje %d", i+1)
		err := channel.Publish(
			"",        // exchange
			queueName, // routing key (nombre de la cola)
			false,     // mandatory
			false,     // immediate
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		if err != nil {
			log.Printf("failed to publish message: %v", err)
		} else {
			log.Printf(" [x] Sent %s", message)
		}
	}
}
