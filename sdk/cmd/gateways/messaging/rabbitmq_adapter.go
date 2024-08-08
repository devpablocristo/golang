package messaging

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func StartConsumer(channel *amqp091.Channel, queueName string) {
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
