package dto

import (
	"github.com/devpablocristo/golang/sdk/examples/analytics/internal/core/analytics/entities"
)

type EventMetrics struct {
	EventMetrics []EventReportRequest `json:"eventMetrics"`
}

type EventReportRequest struct {
	EventID   string `json:"eventId"`
	EventName string `json:"eventName"`
	Visits    int    `json:"visits"`
}

func (ems *EventMetrics) ToDomain() *entities.Report {
	report := &entities.Report{}
	report.Metrics = make(map[string]any)
	for _, eventReportDTO := range ems.EventMetrics {
		report.Metrics[eventReportDTO.EventID] = map[string]any{
			"eventName": eventReportDTO.EventName,
			"visits":    eventReportDTO.Visits,
		}
	}
	return report
}
