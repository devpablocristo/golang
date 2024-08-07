package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func StartProducer() {
	// Conexión a RabbitMQ
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	defer conn.Close()

	// Crear un canal
	ch, err := conn.Channel()
	failOnError(err, "No se pudo abrir el canal")
	defer ch.Close()

	// Declarar una cola
	q, err := ch.QueueDeclare(
		"hello", // nombre de la cola
		false,   // durable
		false,   // auto delete
		false,   // exclusive
		false,   // no-wait
		nil,     // argumentos
	)
	failOnError(err, "No se pudo declarar la cola")

	// Publicar un mensaje
	body := "Hola, RabbitMQ!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key (nombre de la cola)
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "No se pudo publicar el mensaje")
	log.Printf(" [x] Enviado %s", body)
}
