package perconn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/defs"

	"github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
	"github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgreSQL() (ports.Repository, error) {
	r, err := sdkpg.Bootstrap("USERS_DB")
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}
	return &PostgreSQL{repository: r}, nil
}

func (r *PostgreSQL) CreatePerson(ctx context.Context, person *entities.Person) error {
	query := `
		INSERT INTO persons (
			uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP
		)
	`
	_, err := r.repository.Pool().Exec(ctx, query,
		person.UUID, person.Cuil, person.Dni, person.FirstName, person.LastName,
		person.Nationality, person.Email, person.Phone,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			return errors.New("person already exists")
		}
		return fmt.Errorf("error creating person: %w", err)
	}
	return nil
}

func (r *PostgreSQL) UpdatePerson(ctx context.Context, person *entities.Person) error {
	query := `
		UPDATE persons
		SET dni = $1, first_name = $2, last_name = $3, nationality = $4, email = $5, phone = $6, updated_at = CURRENT_TIMESTAMP
		WHERE cuil = $7 AND deleted_at IS NULL
	`
	result, err := r.repository.Pool().Exec(ctx, query,
		person.Dni, person.FirstName, person.LastName, person.Nationality,
		person.Email, person.Phone, person.Cuil,
	)
	if err != nil {
		return fmt.Errorf("error updating person: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("person not found")
	}

	return nil
}

func (r *PostgreSQL) FindPersonByCuil(ctx context.Context, cuil string) (*entities.Person, error) {
	query := `
		SELECT uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at
		FROM persons
		WHERE cuil = $1 AND deleted_at IS NULL
	`
	person := &entities.Person{}
	err := r.repository.Pool().QueryRow(ctx, query, cuil).Scan(
		&person.UUID,
		&person.Cuil,
		&person.Dni,
		&person.FirstName,
		&person.LastName,
		&person.Nationality,
		&person.Email,
		&person.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("person not found")
		}
		return nil, fmt.Errorf("error finding person: %w", err)
	}

	return person, nil
}

func (r *PostgreSQL) FindPersonByUUID(ctx context.Context, uuid string) (*entities.Person, error) {
	query := `
		SELECT uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at
		FROM persons
		WHERE uuid = $1 AND deleted_at IS NULL
	`
	person := &entities.Person{}
	err := r.repository.Pool().QueryRow(ctx, query, uuid).Scan(
		&person.UUID,
		&person.Cuil,
		&person.Dni,
		&person.FirstName,
		&person.LastName,
		&person.Nationality,
		&person.Email,
		&person.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("person not found")
		}
		return nil, fmt.Errorf("error finding person by UUID: %w", err)
	}

	return person, nil
}
