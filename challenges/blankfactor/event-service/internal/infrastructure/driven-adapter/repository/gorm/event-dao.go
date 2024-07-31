package gorm

import (
	"time"

	"github.com/devpablocristo/blankfactor/event-service/internal/domain"
	"gorm.io/gorm"
)

type eventDAO struct {
	gorm.Model
	ID        int64     `gorm:"primary_key"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
}

func (dao *eventDAO) dao2Event() *domain.Event {
	return &domain.Event{
		StartTime: dao.StartTime,
		EndTime:   dao.EndTime,
	}
}

func (dao *eventDAO) TableName() string {
	return "events"
}
