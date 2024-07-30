package ltp

import (
	"context"
	"time"
)

type RepositoryPort interface {
	GetLTP(context.Context, string) error
}

type APIClientPort interface {
	GetKrakenLTPs(context.Context, []string) ([]LTP, time.Time, error)
}
