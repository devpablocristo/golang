package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	hdl "github.com/devpablocristo/qh-users/cmd/api/handlers"
	ucs "github.com/devpablocristo/qh-users/internal/core"
	usr "github.com/devpablocristo/qh-users/internal/core/user"
)

func main() {
	log.Println("Starting application...")

	rabbitmqURI := os.Getenv("RABBITMQ_URI")
	if rabbitmqURI == "" {
		log.Fatal("RABBITMQ_URI no está definido en las variables de entorno")
	}
	log.Println("RABBITMQ_URI:", rabbitmqURI)

	log.Println("Initializing repositories and use cases...")
	repository := usr.NewRepository()
	usecase := ucs.NewUseCase(repository)
	restHandler := hdl.NewRestHandler(usecase)

	log.Println("Initializing RabbitMQ...")
	rabbitMqManager, err := hdl.NewRabbitHandler(rabbitmqURI, "myQueue", usecase)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer func() {
		log.Println("Closing RabbitMQ connection...")
		rabbitMqManager.Close()
	}()

	log.Println("Starting to consume messages...")
	err = rabbitMqManager.ConsumeMessages()
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %s", err)
	}

	r := gin.Default()
	log.Println("Setting up routes...")
	r.GET("/users/:id", restHandler.GetUser)
	r.POST("/users", restHandler.CreateUser)
	r.PUT("/users/:id", restHandler.UpdateUser)
	r.DELETE("/users/:id", restHandler.DeleteUser)
	r.GET("/users", restHandler.ListUsers)
	r.GET("/", restHandler.HelloWorld)

	go func() {
		log.Println("Running server on port 8080...")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %s", err)
		}
	}()

	// Manejar señales del sistema para una finalización ordenada
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
