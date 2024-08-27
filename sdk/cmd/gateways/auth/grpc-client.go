package auth

import (
	"context"
	"fmt"

	pb "github.com/devpablocristo/golang/sdk/pb"
	sdk "github.com/devpablocristo/golang/sdk/pkg/grpc/google/client/ports"
)

type GrpcClient struct {
	client pb.UserServiceClient
	conn   sdk.Client
}

func NewGrpcClient(grpcClient sdk.Client) (*GrpcClient, error) {
	conn, err := grpcClient.GetConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to get gRPC connection: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	return &GrpcClient{
		client: client,
		conn:   grpcClient,
	}, nil
}

func (g *GrpcClient) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	resp, err := g.client.GetUserUUID(ctx, &pb.GetUserRequest{
		Username:     username,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return "", fmt.Errorf("error calling GetUserUUID: %v", err)
	}

	return resp.UUID, nil
}
