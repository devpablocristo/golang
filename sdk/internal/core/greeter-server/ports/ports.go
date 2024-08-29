package ports

import "context"

type UseCases interface {
	SayHello(context.Context) (string, error)
}
