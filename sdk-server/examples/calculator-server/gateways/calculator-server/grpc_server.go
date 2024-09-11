package calculator

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/examples/calculator-server/internal/calculator-server/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkgrpcserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

type CalculatorGrpcServer struct {
	pb.UnimplementedCalculatorServiceServer // NOTE: que es esto?
	useCases                                ports.UseCases
	grpcServer                              sdkgrpcserverport.Server
}

func NewGrpcServer(ucs ports.UseCases, gsv sdkgrpcserverport.Server) *CalculatorGrpcServer {
	return &CalculatorGrpcServer{
		useCases:   ucs,
		grpcServer: gsv,
	}
}

func (s *CalculatorGrpcServer) Start() error {
	s.grpcServer.RegisterService(&pb.CalculatorService_ServiceDesc, s)
	if err := s.grpcServer.Start(); err != nil {
		return err
	}
	return nil
}

func (s *CalculatorGrpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	result, err := s.useCases.Addition(ctx, req.FirstNumber, req.SecondNumber)
	if err != nil {
		return nil, err
	}
	return &pb.SumResponse{SumResult: result}, nil
}
