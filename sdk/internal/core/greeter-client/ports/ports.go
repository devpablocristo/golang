package ports

import "context"

type GrpcClient interface {
	Greet(context.Context) (string, error)
}

type UseCases interface {
	Greet(context.Context) (string, error)
}
