package handler

import (
	"context"
	"fmt"

	ucs "github.com/devpablocristo/qh-users/internal/core"
	"github.com/streadway/amqp"
)

type RabbitHandler struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
	ucs  ucs.UseCasePort
}

// NewRabbitHandler initializes a new RabbitHandler
func NewRabbitHandler(uri, queueName string, ucs ucs.UseCasePort) (*RabbitHandler, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	q, err := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &RabbitHandler{
		conn: conn,
		ch:   ch,
		q:    q,
		ucs:  ucs,
	}, nil
}

func (r *RabbitHandler) PublishMessage(message string) error {
	err := r.ch.Publish(
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // hacer que el mensaje sea persistente
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}
	fmt.Printf("Message published to queue %s: %s\n", r.q.Name, message)
	return nil
}

func (r *RabbitHandler) ConsumeMessages() error {
	messages, err := r.ch.Consume(
		r.q.Name, // nombre de la cola
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %w", err)
	}

	go func() {
		for delivery := range messages {
			fmt.Printf("Message received: %s\n", delivery.Body)
			_, err := r.ucs.PublishMessage(context.Background(), string(delivery.Body))
			if err != nil {
				fmt.Printf("Failed to process message: %s\n", err)
				continue
			}
			fmt.Println("Message acknowledged")
		}
	}()

	return nil
}

// Close closes the channel and connection
func (r *RabbitHandler) Close() {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
