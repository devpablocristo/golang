package authgtw

import (
	"log"

	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkgrpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type GreeterGrpcServer struct {
	pb.UnimplementedGreeterServer // NOTE: que es esto?
	useCases                      ports.UseCases
	grpcServer                    sdkports.Server
}

func NewGrpcServer(ucs ports.UseCases) *GreeterGrpcServer {
	gsv, err := sdkgrpcserver.Bootstrap()
	if err != nil {
		log.Fatalf("failed to initialize gRPC server: %v", err)
	}

	return &GreeterGrpcServer{
		useCases:   ucs,
		grpcServer: gsv,
	}
}

func (s *GreeterGrpcServer) Start() error {
	s.grpcServer.RegisterService(&pb.Greeter_ServiceDesc, s)
	if err := s.grpcServer.Start(); err != nil {
		return err
	}
	return nil
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
