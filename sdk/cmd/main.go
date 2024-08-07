package main

import (
	"fmt"
	"log"

	monitoring "github.com/devpablocristo/golang/sdk/cmd/rest/monitoring/routes"
	user "github.com/devpablocristo/golang/sdk/cmd/rest/user/routes"

	gingonicsetup "github.com/devpablocristo/golang/sdk/internal/platform/gin"
	gmwsetup "github.com/devpablocristo/golang/sdk/internal/platform/go-micro-web"
	initialsetup "github.com/devpablocristo/golang/sdk/internal/platform/initial"
	amqpsetup "github.com/devpablocristo/golang/sdk/internal/platform/rabbitmq"
)

func main() {
	if err := initialsetup.BasicSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	initialsetup.LogInfo("Application started with JWT secret key: %s", initialsetup.GetJWTSecretKey())
	initialsetup.MicroLogInfo("Starting application...")

	//TODO: configurar consul para centrarlizar todas las envs

	gomicro, err := gmwsetup.NewGoMicroInstance()
	if err != nil {
		initialsetup.MicroLogError("error initializing Go Micro: %v", err)
	}

	gingonic, err := gingonicsetup.NewGinInstance()
	if err != nil {
		initialsetup.MicroLogError("error initializing Gin: %v", err)
	}

	monitoring.Routes(gingonic)

	r := gingonic.GetRouter()

	user.Routes(r)

	gomicro.GetService().Handle("/", r)

	// Ejecuta Gin en la dirección especificada por Go-Micro
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run Gin: %v", err)
		}
	}()

	go messaging()

	if err := gomicro.GetService().Run(); err != nil {
		initialsetup.MicroLogError("error starting GoMicro service: %v", err)
	}
}

func messaging() {
	client, err := amqpsetup.NewRabbitMQInstance()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	channel, err := client.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer channel.Close()

	// Define la cola
	queueName := "exampleQueue"
	_, err = channel.QueueDeclare(
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

	// Consume mensajes
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
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	for msg := range messages {
		fmt.Printf("Received message: %s\n", string(msg.Body))
	}
}
