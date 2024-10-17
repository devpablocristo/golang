package ports

import "context"

type UseCases interface {
	Addition(context.Context, int32, int32) (int32, error)
	Subtraction(context.Context, int32, int32) (int32, error)
	Multiplication(context.Context, int32, int32) (int32, error)
	Division(context.Context, int32, int32) (int32, error)
}
