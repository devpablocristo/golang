package user

import (
	"context"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/user/gtwports"
	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb" // Ajusta la ruta según tu proyecto
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
}

func NewGrpcServer() gtwports.GrpcServer {
	return &grpcServer{}
}

func (s *grpcServer) GetUserUUID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	UUID := "userUUID"
	return &pb.GetUserResponse{
		UUID: UUID,
	}, nil
}
