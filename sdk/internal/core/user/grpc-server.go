package user

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	ucs ports.UserUseCases
}

func NewServer(ucs ports.UserUseCases) *Server {
	return &Server{
		ucs: ucs,
	}
}

func (s *Server) GetUserUUID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

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
