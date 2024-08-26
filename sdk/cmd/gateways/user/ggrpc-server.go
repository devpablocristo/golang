package user

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
)

type GgrpcServer struct {
	pb.UnimplementedUserServiceServer
	ucs ports.UserUseCases
}

func NewGgrpcServer(ucs ports.UserUseCases) *GgrpcServer {
	return &GgrpcServer{
		ucs: ucs,
	}
}

func (s *GgrpcServer) GetUserUUID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

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
