package core

import (
	"context"

	ltp "github.com/devpablocristo/dive-challenge/internal/core/ltp"
)

type UseCasePort interface {
	GetLTP(context.Context, []string) ([]ltp.LTP, error)
}
