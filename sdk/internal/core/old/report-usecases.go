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
	report report.Repository
}

func NewUseCases(r report.Repository) UseCasesPort {
	return &UseCases{
		report: r,
	}
}

func (u *UseCases) CreateReport(ctx context.Context, report *report.Report) error {
	fmt.Println(report)
	return nil
}
