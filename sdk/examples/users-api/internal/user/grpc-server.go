package user

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/examples/users-api/internal/user/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
)

type Grpc struct {
	pb.UnimplementedUserServiceServer
	ucs ports.UseCases
}

func NewGrpc(ucs ports.UseCases) ports.GrpcServer {
	return &Grpc{
		ucs: ucs,
	}
}

func (s *Grpc) GetUserUUID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	userUUID, err := s.ucs.GetUserUUID(ctx, req.Username, req.PasswordHash)
	if err != nil {
		return &pb.GetUserResponse{}, err
	}

	_ = userUUID
	UUID := "userUUID"
	return &pb.GetUserResponse{
		UUID: UUID,
	}, nil
}
