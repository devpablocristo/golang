package core

import (
	"context"
	"fmt"

	"github.com/devpablocristo/qh/analytics/internal/core/report"
)

type UseCase struct {
	report report.RepositoryPort
}

func NewUseCase(r report.RepositoryPort) UseCasePort {
	return &UseCase{
		report: r,
	}
}

func (u *UseCase) CreateReport(ctx context.Context, report *report.Report) error {
	fmt.Println(report)
	return nil
}