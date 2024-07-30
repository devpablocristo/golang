package port

import "context"

type ValidateUser interface {
	Execute(context.Context, string, string) (int16, error)
}
