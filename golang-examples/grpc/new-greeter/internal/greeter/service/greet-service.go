package service

import (
	port "github.com/devpablocristo/qh/internal/greeter/ports"
)

// Declare a struct that embeds the GreetServiceServer interface generated from the gRPC proto file
type greetService struct {
}

func NewGreetService(gr port.Repo) port.Service {
	return &greetService{}
}

// Implement the SayHello method of the GreetServiceServer interface
func (s *greetService) SayHello() error {
	return nil
}
