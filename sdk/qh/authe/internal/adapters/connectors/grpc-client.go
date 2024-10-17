package authconn

import (
	"context"
	"fmt"
	"time"

	pb "github.com/devpablocristo/golang/sdk/pb"
	sdk "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-client/ports"
	entities "github.com/devpablocristo/golang/sdk/qh/authe/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/qh/authe/internal/core/ports"
)

type grpcClient struct {
	client     sdkports.Client
	serverName string
}

// NewGrpcClient crea un nuevo cliente gRPC para interactuar con el servicio de usuarios
func NewGrpcClient() (ports.GrpcClient, error) {
	c, err := sdk.Bootstrap() // Usamos tu método Bootstrap del SDK go-micro
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Go Micro gRPC client: %w", err)
	}

	return &grpcClient{
		client: c,
	}, nil
}

func (g *grpcClient) GetClient() sdkports.Client {
	return g.client
}

func (g *grpcClient) GetUserUUID(ctx context.Context, cred *entities.LoginCredentials) (string, error) {
	req := &pb.GetUserRequest{
		Username:     cred.Username,
		PasswordHash: cred.PasswordHash,
	}

	client := g.client.GetClient()
	request := client.NewRequest(g.client.GetServerName(), "GetUserUUID", req)

	var res pb.GetUserResponse

	// Crear un contexto con timeout para la llamada gRPC
	callCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := g.client.GetClient().Call(callCtx, request, &res); err != nil {
		return "", fmt.Errorf("error calling GetUserUUID: %w", err)
	}

	return res.UUID, nil
}
