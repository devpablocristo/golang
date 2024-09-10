package greeter

import (
	"context"
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/internal/core/greeter-client/ports"
	pb "github.com/devpablocristo/golang/sdk/pb"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
)

type grpcClient struct {
	client sdkports.Client
}

func NewGrpcClient(c sdkports.Client) ports.GrpcClient {
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

func (c *grpcClient) GreetClientStream(ctx context.Context, firstName, lastName string) error {
	request := &pb.GreetServerRequest{
		Greeting: &pb.Greeting{
			FirstName: firstName,
			LastName:  lastName,
		},
	}

	// Usar el método genérico InvokeMethod para manejar el stream
	err := c.client.InvokeMethod(ctx, "/greeter.Greeter/GreetServerStreaming", request, func(response any) error {
		// Aquí se procesa cada respuesta del stream
		if res, ok := response.(*pb.GreetServerResponse); ok {
			fmt.Printf("Received streaming message: %s\n", res.Result)
		} else {
			return fmt.Errorf("invalid response type")
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error streaming greet: %v", err)
	}

	return nil
}
