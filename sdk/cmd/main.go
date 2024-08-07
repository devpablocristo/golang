package main

import (
	"log"

	monitoring "github.com/devpablocristo/golang/sdk/cmd/rest/monitoring/routes"
	user "github.com/devpablocristo/golang/sdk/cmd/rest/user/routes"

	msg "github.com/devpablocristo/golang/sdk/cmd/rabbitmq/messaging"
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

	// Configurar y verificar Go Micro
	gomicro, err := gmwsetup.NewGoMicroInstance()
	if err != nil {
		initialsetup.MicroLogError("error initializing Go Micro: %v", err)
	}

	// Configurar y verificar Gin
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

	// Iniciar mensajería (productor y consumidor)
	go messaging()

	// Ejecutar el servicio Go Micro
	if err := gomicro.GetService().Run(); err != nil {
		initialsetup.MicroLogError("error starting GoMicro service: %v", err)
	}
}

// messaging inicializa el productor y consumidor de RabbitMQ
// messaging inicializa el productor y consumidor de RabbitMQ
func messaging() {
	client, err := amqpsetup.NewRabbitMQInstance()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	c, err := client.Channel()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ chan: %v", err)
	}

	// Iniciar consumidor
	go msg.StartConsumer(c, "exampleQueue")

	// Iniciar productor
	go msg.StartProducer(c, "exampleQueue")
}
