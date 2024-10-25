package authconn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
)

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgreSQL() (ports.Repository, error) {
	r, err := sdkpg.Bootstrap("USERS_DB")
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &PostgreSQL{
		repository: r,
	}, nil
}

// Create inserts a new user into the database
func (r *PostgreSQL) CreateUser(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (
			uuid, person_uuid, company_uuid, user_type, 
			username, password_hash, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, 
			$5, $6, 
			CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		)
	`

	_, err := r.repository.Pool().Exec(ctx, query,
		user.UUID,
		user.PersonUUID, // Could be NULL if the user is a company
		user.Credentials.Username,
		user.Credentials.PasswordHash,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			return errors.New("user already exists")
		}
		return err
	}

	return nil
}

// FindByID searches for a user by their UUID
func (r *PostgreSQL) FindUserByUUID(ctx context.Context, id string) (*entities.User, error) {
	user := &entities.User{}
	query := `
		SELECT uuid, person_uuid, company_uuid, user_type, 
			   username, password_hash, created_at, updated_at, deleted_at
		FROM users 
		WHERE uuid = $1 AND deleted_at IS NULL
	`

	err := r.repository.Pool().QueryRow(ctx, query, id).Scan(
		&user.UUID,
		&user.PersonUUID,
		&user.Credentials.Username,
		&user.Credentials.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByCUIL searches for a user by their CUIL (for persons)
func (r *PostgreSQL) FindUserByCuit(ctx context.Context, cuil string) (*entities.User, error) {
	user := &entities.User{}
	query := `
		SELECT u.uuid, u.person_uuid, u.company_uuid, u.user_type, 
		       u.username, u.password_hash, u.created_at, u.updated_at, u.deleted_at
		FROM users u
		JOIN persons p ON u.person_uuid = p.uuid 
		WHERE p.cuil = $1 AND u.deleted_at IS NULL
	`

	err := r.repository.Pool().QueryRow(ctx, query, cuil).Scan(
		&user.UUID,
		&user.PersonUUID,
		&user.Credentials.Username,
		&user.Credentials.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByCUIT searches for a user by their CUIL (for companies)
func (r *PostgreSQL) FindUserByCUIT(ctx context.Context, cuil string) (*entities.User, error) {
	user := &entities.User{}
	query := `
		SELECT u.uuid, u.person_uuid, u.company_uuid, u.user_type, 
		       u.username, u.password_hash, u.created_at, u.updated_at, u.deleted_at
		FROM users u
		JOIN companies c ON u.company_uuid = c.uuid 
		WHERE c.cuil = $1 AND u.deleted_at IS NULL
	`

	err := r.repository.Pool().QueryRow(ctx, query, cuil).Scan(
		&user.UUID,
		&user.PersonUUID,
		&user.Credentials.Username,
		&user.Credentials.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update updates an existing user
func (r *PostgreSQL) UpdateUser(ctx context.Context, user *entities.User) error {
	query := `
		UPDATE users 
		SET person_uuid = $2, company_uuid = $3, user_type = $4, 
			username = $5, password_hash = $6, updated_at = CURRENT_TIMESTAMP
		WHERE uuid = $1 AND deleted_at IS NULL
	`
	result, err := r.repository.Pool().Exec(ctx, query,
		user.UUID,
		user.PersonUUID,
		user.Credentials.Username,
		user.Credentials.PasswordHash,
	)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// SoftDelete marks a user as deleted
func (r *PostgreSQL) SoftDeleteUser(ctx context.Context, id string) error {
	query := `
		UPDATE users 
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE uuid = $1 AND deleted_at IS NULL
	`

	result, err := r.repository.Pool().Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
