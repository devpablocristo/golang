package user

import (
	"context"

	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/user/portsgtw"
	"github.com/devpablocristo/golang/sdk/internal/core/user/portscore"
)

type ggrpcServer struct {
	pb.UnimplementedUserServiceServer
	ucs portscore.UserUseCases
}

func NewGgrpcServer(ucs portscore.UserUseCases) portsgtw.GgrpcServer {
	return &ggrpcServer{
		ucs: ucs,
	}
}

func (s *ggrpcServer) GetUserUUID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

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
