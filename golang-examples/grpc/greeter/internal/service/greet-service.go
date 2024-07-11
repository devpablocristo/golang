package service

import (
	port "greet/internal/service/ports"
)

// Declare a struct that embeds the GreetServiceServer interface generated from the gRPC proto file
type greetService struct {
}

func NewGreetService() port.GreetService {
	return &greetService{}
}

// Implement the SayHello method of the GreetServiceServer interface
func (s *greetService) SayHello() error {
	return nil
}
