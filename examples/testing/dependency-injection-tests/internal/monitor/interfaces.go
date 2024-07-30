package monitor

import (
	"context"
)

type Monitor struct {
	ID          int
	Link        string
	DatadogID   int
	DowntimeID  string
	Product     string
	BrandName   string
	Site        string
	Description string
}

type MonitorRepositoryPort interface {
	CreateMonitor(*Monitor) error
	CheckMonitorExists(*Monitor) (bool, error)
}

type MonitorDatadogPort interface {
	CreateMonitor(ctx context.Context, options ...interface{}) (*Response, error)
}

type Response struct {
	Body []byte
}

type MonitorUsecasePort interface {
	CreateMonitor(ctx context.Context, m *Monitor) error
	PlatformOrBrand(ctx context.Context, m *Monitor, query []string) ([]string, error)
}
