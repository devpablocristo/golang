package core

import (
	"context"

	"github.com/devpablocristo/qh/analytics/internal/core/report"
)

type UseCasePort interface {
	CreateReport(context.Context, *report.Report) error
}
