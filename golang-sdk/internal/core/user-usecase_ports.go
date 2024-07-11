package core

import (
	"context"

	usr "github.com/devpablocristo/qh/events/internal/core/user"
)

type UserUseCasePort interface {
	GetUser(context.Context, string) (usr.User, error)
}
