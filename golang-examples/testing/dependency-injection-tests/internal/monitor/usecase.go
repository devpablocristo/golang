package monitor

import (
	"context"
	"errors"
	"fmt"
)

type MonitorUsecase struct {
	repository          MonitorRepositoryPort
	datadog             MonitorDatadogPort
	PlatformOrBrandFunc func(ctx context.Context, m *Monitor, query []string) ([]string, error)
}

func NewMonitorUsecase(repo MonitorRepositoryPort, datadog MonitorDatadogPort) MonitorUsecasePort {
	return &MonitorUsecase{
		repository: repo,
		datadog:    datadog,
	}
}

func (u *MonitorUsecase) CreateMonitor(ctx context.Context, m *Monitor) error {
	if m == nil {
		return errors.New("monitor cannot be nil")
	}

	exists, _ := u.repository.CheckMonitorExists(m)
	if exists {
		return errors.New("monitor already exists")
	}

	query, err := u.PlatformOrBrand(ctx, m, nil)
	if err != nil {
		return fmt.Errorf("error checking platform or brand: %w", err)
	}

	for _, q := range query {
		r, err := u.datadog.CreateMonitor(ctx, q)
		if err != nil {
			return fmt.Errorf("error creating datadog monitor: %w", err)
		}

		m.DatadogID = int(r.Body[0]) // simplificado para la demo

		err = u.repository.CreateMonitor(m)
		if err != nil {
			return fmt.Errorf("error saving monitor to db: %w", err)
		}
	}
	return nil
}

func (u *MonitorUsecase) PlatformOrBrand(ctx context.Context, m *Monitor, query []string) ([]string, error) {
	if m == nil {
		return nil, errors.New("monitor cannot be nil")
	}

	if u.PlatformOrBrandFunc != nil {
		return u.PlatformOrBrandFunc(ctx, m, query)
	}

	if m.Product == "Platform" {
		query = append(query, "platform_query")
	} else if m.Product == "Brand" {
		query = append(query, "brand_query")
	} else {
		return nil, errors.New("unknown product")
	}
	return query, nil
}
