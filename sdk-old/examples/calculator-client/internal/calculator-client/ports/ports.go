package ports

import "context"

type GrpcClient interface {
	Addition(context.Context, int32, int32) (int32, error)
}

type UseCases interface {
	Addition(context.Context, int32, int32) (int32, error)
}
