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
	r, err := sdkpg.Bootstrap("USERS_DB")
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}
	return &PostgreSQL{repository: r}, nil
}

func (r *PostgreSQL) CreatePerson(ctx context.Context, person *entities.Person) error {
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

func (r *PostgreSQL) UpdatePerson(ctx context.Context, person *entities.Person) error {
	query := `
        UPDATE persons 
        SET dni = $2, first_name = $3, last_name = $4, nationality = $5, email = $6, phone = $7, updated_at = CURRENT_TIMESTAMP
        WHERE cuil = $1 AND deleted_at IS NULL
    `
	result, err := r.repository.Pool().Exec(ctx, query,
		person.Cuil, person.Dni, person.FirstName, person.LastName, person.Nationality,
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

func (r *PostgreSQL) FindPersonByCuil(ctx context.Context, cuil string) (*entities.Person, error) {
	person := &entities.Person{}
	query := `
        SELECT uuid, cuil, dni, first_name, last_name, nationality, email, phone, created_at, updated_at, deleted_at
        FROM persons 
        WHERE cuil = $1 AND deleted_at IS NULL
    `
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return person, nil
}
