package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"demo-project/internal/monitor"
	mock_monitor "demo-project/internal/monitor/mock"
)

func TestCreateMonitor(t *testing.T) {
	type args struct {
		ctx context.Context
		m   *monitor.Monitor
	}
	tests := []struct {
		name            string
		repo            func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort
		datadog         func(ctrl *gomock.Controller) monitor.MonitorDatadogPort
		platformOrBrand func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error)
		args            args
		wantErr         bool
		err             error
	}{
		{
			name: "monitor already exists",
			repo: func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort {
				mockRepo := mock_monitor.NewMockMonitorRepositoryPort(ctrl)
				mockRepo.EXPECT().CheckMonitorExists(gomock.Any()).Return(true, nil)
				return mockRepo
			},
			datadog: func(ctrl *gomock.Controller) monitor.MonitorDatadogPort {
				return nil
			},
			platformOrBrand: func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error) {
				return nil, errors.New("unknown product")
			},
			args: args{
				ctx: context.Background(),
				m: &monitor.Monitor{
					ID:          123,
					Link:        "https://example.com",
					DatadogID:   0,
					DowntimeID:  "456",
					Product:     "Platform",
					BrandName:   "ExampleBrand",
					Site:        "MLA",
					Description: "Monitoring service for ExampleBrand",
				},
			},
			wantErr: true,
			err:     errors.New("monitor already exists"),
		},
		{
			name: "error checking platform or brand",
			repo: func(ctrl *gomock.Controller) monitor.MonitorRepositoryPort {
				mockRepo := mock_monitor.NewMockMonitorRepositoryPort(ctrl)
				mockRepo.EXPECT().CheckMonitorExists(gomock.Any()).Return(false, nil)
				return mockRepo
			},
			datadog: func(ctrl *gomock.Controller) monitor.MonitorDatadogPort {
				mockDatadog := mock_monitor.NewMockMonitorDatadogPort(ctrl)
				// No esperamos que CreateMonitor sea llamado en este caso, ya que debe fallar antes.
				mockDatadog.EXPECT().CreateMonitor(gomock.Any(), gomock.Any()).Times(0)
				return mockDatadog
			},
			platformOrBrand: func(ctx context.Context, m *monitor.Monitor, query []string) ([]string, error) {
				return nil, errors.New("unknown product")
			},
			args: args{
				ctx: context.Background(),
				m: &monitor.Monitor{
					ID:          123,
					Link:        "https://example.com",
					DatadogID:   0,
					DowntimeID:  "456",
					Product:     "Platform",
					BrandName:   "ExampleBrand",
					Site:        "MLA",
					Description: "Monitoring service for ExampleBrand",
				},
			},
			wantErr: true,
			err:     errors.New("error checking platform or brand: unknown product"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := tt.repo(ctrl)
			datadog := tt.datadog(ctrl)

			usecase := monitor.NewMonitorUsecase(repo, datadog)

			// Mock de PlatformOrBrand a nivel de prueba
			usecase.(*monitor.MonitorUsecase).PlatformOrBrandFunc = tt.platformOrBrand

			err := usecase.CreateMonitor(tt.args.ctx, tt.args.m)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMonitor() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
