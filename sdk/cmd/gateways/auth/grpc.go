package auth

import (
	"context"
	pb "path/to/proto/user" // Ajusta la ruta según tu proyecto

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"
	"google.golang.org/grpc"
)

type grpcClient struct{}

func NewGrpcClent() gtwports.GrpcClient {
	return &grpcClient{}
}

func (g *grpcClient) GetUserUUID(username, password string) (string, error) {
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
