package greeter

import (
	"context"
	"fmt"
	"io"
	"time"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
)

// Implementación del servidor Greeter
type GreeterServer struct {
	pb.UnimplementedGreeterServer // NOTE: que es esto?
}

func NewGrpcServer() ports.GrpcServer {
	return &GreeterServer{}
}

// Unary RPC: responde con un mensaje de saludo
func (s *GreeterServer) SayHelloUnary(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloResponse{Message: message}, nil
}

// Server Streaming RPC: envía varios mensajes de saludo al cliente
func (s *GreeterServer) SayHelloServerStreaming(req *pb.HelloRequest, stream pb.Greeter_SayHelloServerStreamingServer) error {
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Hello, %s! Count: %d", req.Name, i)
		if err := stream.Send(&pb.HelloResponse{Message: message}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // Simulación de una operación larga
	}
	return nil
}

// Client Streaming RPC: recibe múltiples solicitudes de saludo y responde con un solo mensaje
func (s *GreeterServer) SayHelloClientStreaming(stream pb.Greeter_SayHelloClientStreamingServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Cuando el cliente ha terminado de enviar solicitudes, enviar una respuesta
			message := fmt.Sprintf("Hello, %s!", names)
			return stream.SendAndClose(&pb.HelloResponse{Message: message})
		}
		if err != nil {
			return err
		}
		names = append(names, req.Name)
	}
}

// Bidirectional Streaming RPC: el servidor y el cliente envían un flujo de mensajes
func (s *GreeterServer) SayHelloBidirectionalStreaming(stream pb.Greeter_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %s!", req.Name)
		if err := stream.Send(&pb.HelloResponse{Message: message}); err != nil {
			return err
		}
	}
	return nil
}
