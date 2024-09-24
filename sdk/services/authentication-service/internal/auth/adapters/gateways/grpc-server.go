package authgtw

import (
	"fmt"

	pb "github.com/devpablocristo/golang/sdk/pb"
	sdk "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-server/ports"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type GreeterGrpcServer struct {
	pb.UnimplementedGreeterServer // NOTE: que es esto?
	useCases                      ports.UseCases
	grpcServer                    sdkports.Server
}

func NewGrpcServer(ucs ports.UseCases) (*GreeterGrpcServer, error) {
	gsv, err := sdk.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gRPC server: %w", err)
	}

	return &GreeterGrpcServer{
		useCases:   ucs,
		grpcServer: gsv,
	}, nil
}

func (s *GreeterGrpcServer) GetServer() sdkports.Server {
	return s.grpcServer
}

// func (s *GreeterGrpcServer) GreetUnary(ctx context.Context, req *pb.GreetUnaryRequest) (*pb.GreetUnaryResponse, error) {
// 	name := req.GetGreeting().GetFirstName() + " " + req.GetGreeting().GetLastName()
// 	greeting, err := s.useCases.Greet(ctx, name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GreetUnaryResponse{Result: greeting}, nil
// }

// func (s *GreeterGrpcServer) GreetServerStreaming(req *pb.GreetServerRequest, stream pb.Greeter_GreetServerStreamingServer) error {
// 	for i := 0; i < 5; i++ {
// 		message := fmt.Sprintf("Hello, %s %s! Count: %d", req.Greeting.FirstName, req.Greeting.LastName, i)
// 		if err := stream.Send(&pb.GreetServerResponse{Result: message}); err != nil {
// 			return err
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// 	return nil
// }
