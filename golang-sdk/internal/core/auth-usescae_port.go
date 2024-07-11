package core

import (
	"context"
)

type AuthUseCasePort interface {
	Login(context.Context, string, string) (string, error)
}
