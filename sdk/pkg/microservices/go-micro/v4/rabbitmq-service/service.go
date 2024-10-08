package main

import (
	"fmt"
	"log"
	"time"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"

	ports "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/prod-cons/ports"
)

func main() {

	var config ports.Config

	// Configurar el broker para usar RabbitMQ
	rabbit := broker.NewBroker(
		broker.Addrs(config.GetAddress()),
	)

	// Inicializar el broker
	if err := rabbit.Init(); err != nil {
		log.Fatalf("Error al inicializar el broker: %v", err)
	}

	if err := rabbit.Connect(); err != nil {
		log.Fatalf("Error al conectar con el broker: %v", err)
	}

	// Crear el servicio
	service := micro.NewService(
		micro.Name("producer-consumer"),
		micro.Broker(rabbit),
	)

	service.Init()

	// Función para consumir mensajes
	go func() {
		// Suscribirse al tema "topic.hola"
		if _, err := rabbit.Subscribe("topic.hola", func(p broker.Event) error {
			msg := p.Message()
			fmt.Printf("Mensaje recibido: %s\n", string(msg.Body))
			return nil
		}, broker.Queue("consumer-queue")); err != nil {
			log.Fatalf("Error al suscribirse: %v", err)
		}

		// Mantener el consumidor en ejecución
		if err := service.Run(); err != nil {
			log.Fatalf("Error al ejecutar el servicio: %v", err)
		}
	}()

	// Función para publicar mensajes cada 2 segundos
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		msg := &broker.Message{
			Header: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: []byte("¡Hola desde el producer-consumer!"),
		}

		if err := rabbit.Publish("topic.hola", msg); err != nil {
			fmt.Println("Error al publicar mensaje:", err)
		} else {
			fmt.Println("Mensaje publicado")
		}
	}
}
