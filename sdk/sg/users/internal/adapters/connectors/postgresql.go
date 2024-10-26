package authconn

import (
	"context"
	"errors"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"
	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"

	datmod "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors/data-model"
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

// Función para hashear la contraseña antes de almacenar en la base de datos
func (r *PostgreSQL) CreateUser(ctx context.Context, user *datmod.User) error {
	passwordHash, err := sdktools.HashPassword(user.Password, 12)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	// Consulta SQL para insertar el usuario
	query := `
		INSERT INTO users (
			uuid, person_uuid, email_validated, username, password_hash, created_at
		) VALUES (
			$1, $2, $3, $4, $5, CURRENT_TIMESTAMP
		)
	`

	// Ejecutar la consulta con el hash generado
	_, err = r.repository.Pool().Exec(ctx, query,
		user.UUID,
		user.PersonUUID,
		user.EmailValidated,
		user.Username,
		user.Password, // Aquí se usa el hash de la contraseña
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			return errors.New("user already exists")
		}
		return err
	}

	return nil
}

// // FindByCUIL searches for a user by their CUIL (for persons)
// func (r *PostgreSQL) FindUserByCuil(ctx context.Context, cuil string) (*entities.User, error) {
// 	user := &entities.User{}
// 	query := `
// 		SELECT u.uuid, u.person_uuid, u.username, u.password_hash, u.created_at, u.updated_at, u.deleted_at
// 		FROM users u
// 		JOIN persons p ON u.person_uuid = p.uuid
// 		WHERE p.cuil = $1 AND u.deleted_at IS NULL
// 	`

// 	err := r.repository.Pool().QueryRow(ctx, query, cuil).Scan(
// 		&user.UUID,
// 		&user.PersonUUID,
// 		&user.Credentials.Username,
// 		&user.Credentials.PasswordHash,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 		&user.DeletedAt,
// 	)

// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// // Update updates an existing user
// func (r *PostgreSQL) UpdateUser(ctx context.Context, user *datmod.User) error {
// 	query := `
// 		UPDATE users
// 		SET person_uuid = $2, user_type = $3, username = $4, password_hash = $5, updated_at = CURRENT_TIMESTAMP
// 		WHERE uuid = $1 AND deleted_at IS NULL
// 	`
// 	result, err := r.repository.Pool().Exec(ctx, query,
// 		user.UUID,
// 		user.PersonUUID,
// 		user.Credentials.Username,
// 		user.Credentials.PasswordHash,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("user not found")
// 	}

// 	return nil
// }
