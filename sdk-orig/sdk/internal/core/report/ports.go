package report

import (
	"context"
)

type RepositoryPort interface {
	CreateReport(context.Context) error
}
