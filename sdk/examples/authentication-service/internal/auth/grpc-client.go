package auth

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/examples/authentication-service/internal/auth/entities"
	ports "github.com/devpablocristo/golang/sdk/examples/authentication-service/internal/auth/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
)

type grpcClient struct {
	client sdkports.Client
}

// NewGrpcClient crea un nuevo cliente gRPC para interactuar con el servicio de usuarios
func NewGrpcClient(c sdkports.Client) ports.GrpcClient {
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
