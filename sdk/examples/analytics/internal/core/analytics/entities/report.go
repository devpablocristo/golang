package entities

import "time"

type Report struct {
	ReportID    string
	GeneratedAt time.Time
	Metrics     map[string]any
}

type InMemDB map[string]*Report
