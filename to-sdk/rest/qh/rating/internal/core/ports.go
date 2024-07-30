package core

import (
	"context"

	ltp "github.com/devpablocristo/qh/rating/internal/core/ltp"
)

type UseCasePort interface {
	GetLTP(context.Context, []string) ([]ltp.LTP, error)
}
