package core

import (
	"context"
	"fmt"

	"github.com/devpablocristo/golang/sdk/internal/core/report"
)

type UseCasesPort interface {
	CreateReport(context.Context, *report.Report) error
}

type UseCases struct {
	report report.RepositoryPort
}

func NewUseCases(r report.RepositoryPort) UseCasesPort {
	return &UseCases{
		report: r,
	}
}

func (u *UseCases) CreateReport(ctx context.Context, report *report.Report) error {
	fmt.Println(report)
	return nil
}
