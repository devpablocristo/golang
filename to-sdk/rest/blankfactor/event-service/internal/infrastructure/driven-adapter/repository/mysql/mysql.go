package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	port "github.com/devpablocristo/blankfactor/event-service/internal/application/port"
	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type EventMysqlRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) port.EventRepo {
	return &EventMysqlRepository{
		db: db,
	}
}

func (r *EventMysqlRepository) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	query := "INSERT INTO events (start_time, end_time) VALUES (?, ?)"

	var e eventDAO

	e.StartTime = event.StartTime
	e.EndTime = event.EndTime

	result, err := r.db.ExecContext(ctx, query, e.StartTime, e.EndTime)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error saving event: %v", err)
	}

	fmt.Println(id)

	savedEvent, err := r.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error saving event: %v", err)
	}

	return savedEvent, nil
}

func (r *EventMysqlRepository) GetEventByID(ctx context.Context, id int64) (*domain.Event, error) {
	query := "SELECT id, start_time, end_time FROM events WHERE id = ?"

	row := r.db.QueryRowContext(ctx, query, id)

	e := eventDAO{}
	var startTimeStr, endTimeStr string
	err := row.Scan(&e.ID, &startTimeStr, &endTimeStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	e.StartTime, err = strToTime(startTimeStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	e.EndTime, err = strToTime(endTimeStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	event := e.dao2Event()

	return event, nil
}

func strToTime(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Println(err)
		return time.Time{}, err
	}
	return t, nil
}

func (r *EventMysqlRepository) UpdateEvent(ctx context.Context, event *domain.Event, id int64) error {
	query := "UPDATE events SET start_time = ?, end_time = ? WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, event.StartTime, event.EndTime, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventMysqlRepository) DeleteEvent(ctx context.Context, id int64) error {
	query := "DELETE FROM events WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventMysqlRepository) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	query := "SELECT id, start_time, end_time FROM events"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.New("rows - " + err.Error())
	}
	defer rows.Close()

	events := make([]*domain.Event, 0)

	var e eventDAO
	var startTimeStr, endTimeStr string
	for rows.Next() {
		if err := rows.Scan(&e.ID, &startTimeStr, &endTimeStr); err != nil {
			return nil, errors.New("error de tipo - " + err.Error())
		}
		e.StartTime, err = strToTime(startTimeStr)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		e.EndTime, err = strToTime(endTimeStr)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		event := e.dao2Event()
		events = append(events, event)
	}

	return events, nil
}
