package ports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/examples/analytics/internal/core/analytics/entities"
)

type Repository interface {
	CreateReport(context.Context) error
}

type UseCases interface {
	CreateReport(context.Context, *entities.Report) error
}
