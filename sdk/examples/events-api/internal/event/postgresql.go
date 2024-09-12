package event

import (
	"context"
	"fmt"

	entities "github.com/devpablocristo/golang/sdk/examples/events-api/internal/event/entities"
	ports "github.com/devpablocristo/golang/sdk/examples/events-api/internal/event/ports"
	sdkpostgresqlports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"
)

type repository struct {
	pgInst sdkpostgresqlports.Repository
}

func Newrepository(inst sdkpostgresqlports.Repository) ports.Repository {
	return &repository{
		pgInst: inst,
	}
}

func (r *repository) CreateEvent(ctx context.Context, event *entities.Event) error {
	// query := `INSERT INTO events (name, date) VALUES ($1, $2) RETURNING id`
	// err := r.db.Pool().QueryRow(ctx, query, event.Name, event.Date).Scan(&event.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to create event: %w", err)
	// }

	fmt.Printf("Event created: %+v\n", event)
	return nil
}
