package core

import (
	"context"
	"fmt"
	"time"

	ltp "github.com/devpablocristo/dive-challenge/internal/core/ltp"
)

type UseCase struct {
	ltp ltp.RepositoryPort
	acl ltp.APIClientPort
}

func NewUseCase(r ltp.RepositoryPort, a ltp.APIClientPort) UseCasePort {
	return &UseCase{
		ltp: r,
		acl: a,
	}
}

func (u *UseCase) GetLTP(ctx context.Context, pairs []string) ([]ltp.LTP, error) {
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
