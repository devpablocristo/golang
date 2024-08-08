package report

import "time"

type Report struct {
	ReportID    string
	GeneratedAt time.Time
	Metrics     map[string]interface{}
}

type InMemDB map[string]*Report
