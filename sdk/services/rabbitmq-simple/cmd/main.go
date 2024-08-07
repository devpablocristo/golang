package main

import (
	"github.com/devpablocristo/golang/sdk/services/rabbitmq-simple/cmd/messaging"
)

func main() {
	// Start the producer
	go messaging.StartProducer()

	// Start the consumer
	messaging.StartConsumer()
}
