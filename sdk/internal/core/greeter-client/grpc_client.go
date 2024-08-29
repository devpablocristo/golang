package greeter

import (
	"context"
	"time"

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

func (c *grpcClient) Greet(ctx context.Context, name string) (string, error) {
	// Crear un contexto con un tiempo de espera
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// // Crear una solicitud de tipo HelloRequest
	// request := &pb.GreetUnaryRequest{
	// 	firstName: name,
	// }

	// // Crear una variable para almacenar la respuesta de tipo HelloResponse
	var response pb.GreetUnaryResponse

	// // Invocar el método GreetUnary del servicio Greeter
	// err := c.client.InvokeMethod(ctx, "/greeter.Greeter/GreetUnary", request, &response)
	// if err != nil {
	// 	return "", err // Devolver el error para que sea manejado por el llamador
	// }

	// Devolver el mensaje de la respuesta
	return response.Result, nil
}
