package ports

import "context"

type GrpcClient interface {
	SayHello(context.Context, string) (string, error)
}

type UseCases interface {
	SayHello(context.Context) (string, error)
}
