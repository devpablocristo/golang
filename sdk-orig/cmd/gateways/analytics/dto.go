package handler

import (
	"github.com/devpablocristo/golang/sdk/internal/core/report"
)

type EventMetricsDTO struct {
	EventMetrics []EventReportDTO `json:"eventMetrics"`
}

type EventReportDTO struct {
	EventID   string `json:"eventId"`
	EventName string `json:"eventName"`
	Visits    int    `json:"visits"`
}

func (ems *EventMetricsDTO) ToDomain() *report.Report {
	report := &report.Report{}
	report.Metrics = make(map[string]any)
	for _, eventReportDTO := range ems.EventMetrics {
		report.Metrics[eventReportDTO.EventID] = map[string]any{
			"eventName": eventReportDTO.EventName,
			"visits":    eventReportDTO.Visits,
		}
	}
	return report
}
