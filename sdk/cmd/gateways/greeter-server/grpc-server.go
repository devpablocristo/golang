package greeter

import (
	"context"
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter-server/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkgrpcserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

type GreeterGrpcServer struct {
	pb.UnimplementedGreeterServer // NOTE: que es esto?
	useCases                      ports.UseCases
	grpcServer                    sdkgrpcserverport.Server
}

func NewGrpcServer(ucs ports.UseCases, gsv sdkgrpcserverport.Server) *GreeterGrpcServer {
	return &GreeterGrpcServer{
		useCases:   ucs,
		grpcServer: gsv,
	}
}

func (s *GreeterGrpcServer) Start() error {
	s.grpcServer.RegisterService(&pb.Greeter_ServiceDesc, s.grpcServer)
	if err := s.grpcServer.Start(); err != nil {
		return err
	}
	return nil
}

// Unary RPC: responde con un mensaje de saludo
func (s *GreeterGrpcServer) SayHelloUnary(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {

	message := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.SayHelloResponse{Message: message}, nil
}

// Server Streaming RPC: envía varios mensajes de saludo al cliente
// func (s *GreeterGrpcServer) SayHelloServerStreaming(req *pb.SayHelloRequest, stream pb.Greeter_SayHelloServerStreamingServer) error {
// 	for i := 0; i < 5; i++ {
// 		message := fmt.Sprintf("Hello, %s! Count: %d", req.Name, i)
// 		if err := stream.Send(&pb.SayHelloResponse{Message: message}); err != nil {
// 			return err
// 		}
// 		time.Sleep(1 * time.Second) // Simulación de una operación larga
// 	}
// 	return nil
// }

// // Client Streaming RPC: recibe múltiples solicitudes de saludo y responde con un solo mensaje
// func (s *GreeterGrpcServer) SayHelloClientStreaming(stream pb.Greeter_SayHelloClientStreamingServer) error {
// 	var names []string
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			// Cuando el cliente ha terminado de enviar solicitudes, enviar una respuesta
// 			message := fmt.Sprintf("Hello, %s!", names)
// 			return stream.SendAndClose(&pb.SayHelloResponse{Message: message})
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		names = append(names, req.Name)
// 	}
// }

// // Bidirectional Streaming RPC: el servidor y el cliente envían un flujo de mensajes
// func (s *GreeterGrpcServer) SayHelloBidirectionalStreaming(stream pb.Greeter_SayHelloBidirectionalStreamingServer) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		message := fmt.Sprintf("Hello, %s!", req.Name)
// 		if err := stream.Send(&pb.SayHelloResponse{Message: message}); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
