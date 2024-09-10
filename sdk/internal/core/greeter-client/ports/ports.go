package ports

import "context"

type GrpcClient interface {
	Greet(context.Context, string, string) (string, error)
	GreetClientStream(context.Context, string, string) error
}

type UseCases interface {
	Greet(context.Context, string, string) (string, error)
}
