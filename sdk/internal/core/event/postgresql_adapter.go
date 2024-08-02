package event

import (
	"context"
	"fmt"

	pgxpostgres "github.com/devpablocristo/golang-sdk/pkg/postgresql/pgxpool"
)

// type Repository struct {
// 	inMemDB *InMemDB
// 	db      *pgxpool.Pool
// }

// func NewRepository(db *db.PostgreSQL) RepositoryPort {
// 	inmem := make(InMemDB)
// 	return &Repository{
// 		inMemDB: &inmem,
// 		db:      db.Pool(), // Accediendo al pool de conexiones
// 	}
// }

// func (r *Repository) CreateEvent(ctx context.Context, event *Event) error {
// 	// query := `INSERT INTO events (name, date) VALUES ($1, $2) RETURNING id`
// 	// err := r.db.QueryRow(ctx, query, event.Name, event.Date).Scan(&event.ID)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to create event: %w", err)
// 	// }

// 	fmt.Printf("Event created: %+v\n", event)
// 	return nil
// }

type Repository struct {
	pgInst pgxpostgres.PostgreSQLClientPort
}

func NewRepository(inst pgxpostgres.PostgreSQLClientPort) RepositoryPort {
	return &Repository{
		pgInst: inst,
	}
}

func (r *Repository) CreateEvent(ctx context.Context, event *Event) error {
	// query := `INSERT INTO events (name, date) VALUES ($1, $2) RETURNING id`
	// err := r.db.Pool().QueryRow(ctx, query, event.Name, event.Date).Scan(&event.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to create event: %w", err)
	// }

	fmt.Printf("Event created: %+v\n", event)
	return nil
}
