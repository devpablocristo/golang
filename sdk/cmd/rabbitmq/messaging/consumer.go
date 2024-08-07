package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func StartConsumer() {
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

	// Consumir mensajes
	msgs, err := ch.Consume(
		q.Name, // nombre de la cola
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // argumentos
	)
	failOnError(err, "No se pudo registrar el consumidor")

	forever := make(chan bool)

	// Goroutine para manejar los mensajes
	go func() {
		for d := range msgs {
			log.Printf(" [x] Recibido %s", d.Body)
		}
	}()

	log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")
	<-forever
}
