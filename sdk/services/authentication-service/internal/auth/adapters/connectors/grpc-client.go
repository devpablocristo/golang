package authconn

import (
	"context"
	"log"

	pb "github.com/devpablocristo/golang/sdk/pb"
	sdk "github.com/devpablocristo/golang/sdk/pkg/grpc/client"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	entities "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/entities"
	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type grpcClient struct {
	client sdkports.Client
}

// NewGrpcClient crea un nuevo cliente gRPC para interactuar con el servicio de usuarios
func NewGrpcClient() ports.GrpcClient {
	c, err := sdk.Bootstrap()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return &grpcClient{
		client: c,
	}
}

// GetUserUUID obtiene el UUID del usuario desde el servicio de usuarios
func (g *grpcClient) GetUserUUID(ctx context.Context, cred *entities.LoginCredentials) (string, error) {
	req := &pb.GetUserRequest{
		Username:     cred.Username,
		PasswordHash: cred.PasswordHash,
	}

	var res pb.GetUserResponse

	err := g.client.InvokeMethod(ctx, "/user.UserService/GetUserUUID", req, &res)
	if err != nil {
		return "", err
	}

	return res.UUID, nil
}
