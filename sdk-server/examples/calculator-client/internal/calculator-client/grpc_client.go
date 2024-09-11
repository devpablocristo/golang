package calculator

import (
	"context"

	ports "github.com/devpablocristo/golang/sdk/examples/calculator-client/internal/calculator-client/ports"
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

func (c *grpcClient) Addition(ctx context.Context, firstNum, secondNum int32) (int32, error) {
	request := &pb.SumRequest{
		FirstNumber:  firstNum,
		SecondNumber: secondNum,
	}

	var response pb.SumResponse

	err := c.client.InvokeMethod(ctx, "/calculator.CalculatorService/Sum", request, &response)
	if err != nil {
		return 0, err
	}

	return response.SumResult, nil
}
