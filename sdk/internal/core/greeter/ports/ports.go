package ports

import "context"

type GrpcClient interface {
	SayHello(context.Context, string) (string, error)
}

type UseCases interface {
	Hello(context.Context) (string, error)
}
