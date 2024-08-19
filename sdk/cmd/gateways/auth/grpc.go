package auth

import (
	"context"

	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/pb"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/portsgtw"
	"github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
	"google.golang.org/grpc"
)

type ggrpcClient struct {
	client portsgtw.GrpcClient
}

func NewGGrpcClient(client portspkg.GgrpcClient) portsgtw.GrpcClient {
	return &ggrpcClient{
		client: client,
	}
}

func (g *ggrpcClient) GetUserUUID(username, password string) (string, error) {
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
