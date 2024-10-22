package perconn

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
	ports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgreSQL() (ports.Repository, error) {
	r, err := sdkpg.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}
	return &PostgreSQL{repository: r}, nil
}

// Create inserts a new person into the database
func (r *PostgreSQL) Create(ctx context.Context, person *entities.Person) error {
	query := `
        INSERT INTO persons (
            uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
        )
    `
	_, err := r.repository.Pool().Exec(ctx, query,
		person.UUID, person.Cuil, person.Dni, person.FirstName, person.LastName,
		person.Nationality, person.Email, person.Phone,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			return fmt.Errorf("person already exists")
		}
		return err
	}
	return nil
}

// FindByID searches for a person by their UUID
func (r *PostgreSQL) FindByID(ctx context.Context, UUID string) (*entities.Person, error) {
	person := &entities.Person{}
	query := `
        SELECT uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at, deleted_at
        FROM persons 
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	err := r.repository.Pool().QueryRow(ctx, query, UUID).Scan(
		&person.UUID,
		&person.Cuil,
		&person.Dni,
		&person.FirstName,
		&person.LastName,
		&person.Nationality,
		&person.Email,
		&person.Phone,
		&person.CreatedAt,
		&person.UpdatedAt,
		&person.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return person, nil
}

// Update updates an existing person
func (r *PostgreSQL) Update(ctx context.Context, person *entities.Person) error {
	query := `
        UPDATE persons 
        SET cuil = $2, dni = $3, first_name = $4, last_name = $5, nationality = $6, email = $7, phone = $8, updated_at = CURRENT_TIMESTAMP
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	result, err := r.repository.Pool().Exec(ctx, query,
		person.UUID, person.Cuil, person.Dni, person.FirstName, person.LastName, person.Nationality,
		person.Email, person.Phone,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("person not found")
	}

	return nil
}

// SoftDelete marks a person as deleted
func (r *PostgreSQL) SoftDelete(ctx context.Context, UUID string) error {
	query := `
        UPDATE persons 
        SET deleted_at = CURRENT_TIMESTAMP
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	result, err := r.repository.Pool().Exec(ctx, query, UUID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("person not found")
	}

	return nil
}

// FindByCuit searches for a person by their CUIL
func (r *PostgreSQL) FindByCuit(ctx context.Context, cuit string) (*entities.Person, error) {
	person := &entities.Person{}
	query := `
        SELECT uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at, deleted_at
        FROM persons 
        WHERE cuil = $1 AND deleted_at IS NULL
    `
	err := r.repository.Pool().QueryRow(ctx, query, cuit).Scan(
		&person.UUID,
		&person.Cuil,
		&person.Dni,
		&person.FirstName,
		&person.LastName,
		&person.Nationality,
		&person.Email,
		&person.Phone,
		&person.CreatedAt,
		&person.UpdatedAt,
		&person.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return person, nil
}
