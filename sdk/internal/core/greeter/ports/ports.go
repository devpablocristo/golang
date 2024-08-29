package ports

import "context"

type GrpcServer interface{}

type UseCases interface {
	Hello(context.Context) (string, error)
}
