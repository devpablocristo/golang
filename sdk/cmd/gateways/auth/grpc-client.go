package auth

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/pb"
	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
)

type GrpcClient struct {
	ucs ports.UserUseCases
}

func NewGrpcClient(u ports.UserUseCases) *GrpcClient {
	return &GrpcClient{
		ucs: u,
	}
}

func (g *GrpcClient) GetUserUUID(username, password string) (string, error) {
	conn, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Hacer la solicitud al servicio user
	resp, err := client.GetUserUUID(context.Background(), &pb.GetUserRequest{
		Username:     username,
		PasswordHash: password, // Aquí debería ir el hash de la contraseña
	})
	if err != nil {
		return "", err
	}

	return resp.UUID, nil
}
