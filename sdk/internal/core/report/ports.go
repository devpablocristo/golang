package report

import (
	"context"
)

type Repository interface {
	CreateReport(context.Context) error
}
