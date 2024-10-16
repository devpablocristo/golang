package report

import (
	"context"
	"fmt"

	"github.com/devpablocristo/golang/sdk/services/analytics/internal/analytics/entities"
	"github.com/devpablocristo/golang/sdk/services/analytics/internal/analytics/ports"
)

type UseCases struct {
	report ports.Repository
}

func NewUseCases(r ports.Repository) ports.UseCases {
	return &UseCases{
		report: r,
	}
}

func (u *UseCases) CreateReport(ctx context.Context, report *entities.Report) error {
	fmt.Println(report)
	return nil
}
