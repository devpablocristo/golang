package ports

import "context"

type UseCases interface {
	Greet(context.Context, string) (string, error)
}
