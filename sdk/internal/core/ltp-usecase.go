package core

import (
	"context"
	"fmt"
	"time"

	ltp "github.com/devpablocristo/golang-sdk/internal/core/ltp"
)

type LtpUseCasePort interface {
	GetLTP(context.Context, []string) ([]ltp.LTP, error)
}

type LtpUseCase struct {
	ltp ltp.RepositoryPort
	acl ltp.APIClientPort
}

func NewLtpUseCase(r ltp.RepositoryPort, a ltp.APIClientPort) LtpUseCasePort {
	return &LtpUseCase{
		ltp: r,
		acl: a,
	}
}

func (u *LtpUseCase) GetLTP(ctx context.Context, pairs []string) ([]ltp.LTP, error) {
	lp, timestamp, err := u.acl.GetKrakenLTPs(ctx, pairs)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	timeDiff := currentTime.Sub(timestamp)
	if timeDiff.Minutes() > 1 {
		return nil, fmt.Errorf("invalid data: older than one minute")
	}

	return lp, nil
}
