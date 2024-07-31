package mysql

import (
	"time"

	"github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type eventDAO struct {
	ID        int64     `db:"id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (dao *eventDAO) dao2Event() *domain.Event {
	return &domain.Event{
		StartTime: dao.StartTime,
		EndTime:   dao.EndTime,
	}
}
