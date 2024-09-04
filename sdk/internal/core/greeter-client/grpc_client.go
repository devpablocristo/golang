package greeter

import (
	"context"

	coreports "github.com/devpablocristo/golang/sdk/internal/core/greeter-client/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
)

type grpcClient struct {
	client sdkports.Client
}

func NewGrpcClient(c sdkports.Client) coreports.GrpcClient {
	return &grpcClient{
		client: c,
	}
}

func (c *grpcClient) Greet(ctx context.Context, fistName, lastName string) (string, error) {
	// Crear una solicitud de tipo HelloRequest
	request := &pb.GreetUnaryRequest{
		Greeting: &pb.Greeting{
			FirstName: fistName,
			LastName:  lastName,
		},
	}

	// Crear una variable para almacenar la respuesta de tipo HelloResponse
	var response pb.GreetUnaryResponse

	// Invocar el método GreetUnary del servicio Greeter
	err := c.client.InvokeMethod(ctx, "/greeter.Greeter/GreetUnary", request, &response)
	if err != nil {
		return "", err // Devolver el error para que sea manejado por el llamador
	}

	// Devolver el mensaje de la respuesta
	return response.Result, nil
}
