package authconn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"
	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"

	datmod "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors/data-model"
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



// Crear usuario en la base de datos con hash de contraseña
func (r *PostgreSQL) CreateUser(ctx context.Context, user *datmod.User) error {
	passwordHash, err := sdktools.HashPassword(user.Password, 12)
	if err != nil {
		return err
	}
	user.Password = passwordHash

	query := `
		INSERT INTO users (
			uuid, person_uuid, email_validated, username, password_hash, created_at
		) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
	`

	_, err = r.repository.Pool().Exec(ctx, query,
		user.UUID,
		user.PersonUUID,
		user.EmailValidated,
		user.Username,
		user.Password, // Hash de la contraseña
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // Código para unique_violation
			return errors.New("user already exists")
		}
		return err
	}

	return nil
}

// Buscar usuario por UUID
func (r *PostgreSQL) FindUserByUserUUID(ctx context.Context, userUUID string) (*entities.User, error) {
	query := `SELECT uuid, person_uuid, username, password_hash, email_validated, created_at, updated_at, deleted_at FROM users WHERE uuid = $1`

	return r.findUserByQuery(ctx, query, userUUID)
}

// Actualizar usuario por UUID
func (r *PostgreSQL) UpdateUser(ctx context.Context, user *datmod.User) error {
	query := `
		UPDATE users
		SET email_validated = $1, username = $2, password_hash = $3, updated_at = CURRENT_TIMESTAMP
		WHERE uuid = $4
	`

	_, err := r.repository.Pool().Exec(ctx, query,
		user.EmailValidated,
		user.Username,
		user.Password,
		user.UUID,
	)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

// Buscar usuario por PersonUUID
func (r *PostgreSQL) FindUserByPersonUUID(ctx context.Context, personUUID string) (*entities.User, error) {
	query := `SELECT uuid, person_uuid, username, password_hash, email_validated, created_at, updated_at, deleted_at FROM users WHERE person_uuid = $1`

	return r.findUserByQuery(ctx, query, personUUID)
}

// Función auxiliar para encontrar un usuario usando una consulta SQL
func (r *PostgreSQL) findUserByQuery(ctx context.Context, query string, identifier string) (*entities.User, error) {
	// Ejecutar la consulta
	row := r.repository.Pool().QueryRow(ctx, query, identifier)

	// Mapear el resultado a una entidad User
	var user entities.User
	var personUUID sql.NullString

	err := row.Scan(&user.UUID, &personUUID, &user.Credentials.Username, &user.Credentials.Password, &user.EmailValidated, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found for identifier: %s", identifier)
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	user.PersonUUID = getStringFromNullString(personUUID)

	return &user, nil
}
