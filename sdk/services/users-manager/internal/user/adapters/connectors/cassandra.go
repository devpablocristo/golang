package userconn

import (
	"context"
	"errors"

	"github.com/gocql/gocql"

	sdkports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
	entities "github.com/devpablocristo/golang/sdk/services/users-manager/internal/user/core/entities"
	ports "github.com/devpablocristo/golang/sdk/services/users-manager/internal/user/core/ports"
)

type cassandra struct {
	repository sdkports.Repository
}

func NewCassandraRepository(r sdkports.Repository) ports.Repository {
	return &cassandra{
		repository: r,
	}
}

func (r *cassandra) SaveUser(ctx context.Context, user *entities.User) error {
	return r.repository.GetSession().Query(
		"INSERT INTO users (uuid, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.UUID, user.Credentials.Username, user.Credentials.PasswordHash, user.CreatedAt,
	).Exec()
}

func (r *cassandra) GetUserByUUID(ctx context.Context, UUID string) (*entities.User, error) {
	var user entities.User
	err := r.repository.GetSession().Query(
		"SELECT UUID, username, password, created_at FROM users WHERE uuid = ?",
		UUID,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *cassandra) GetUserByCredentials(ctx context.Context, username, passwordHash string) (string, error) {
	var uuUUID string
	err := r.repository.GetSession().Query(
		"SELECT UUID FROM users WHERE username = ? AND password = ?",
		username, passwordHash,
	).Consistency(gocql.One).Scan(&uuUUID)
	if err != nil {
		return "", err
	}
	return uuUUID, nil
}

func (r *cassandra) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := r.repository.GetSession().Query(
		"SELECT UUID, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *cassandra) DeleteUser(ctx context.Context, UUID string) error {
	return r.repository.GetSession().Query(
		"DELETE FROM users WHERE UUID = ?",
		UUID,
	).Exec()
}

func (r *cassandra) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	userDB := make(entities.InMemDB)

	iter := r.repository.GetSession().Query("SELECT UUID, username, password, created_at FROM users").Iter()

	var user entities.User
	for iter.Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt) {
		// Agregar el usuario al mapa usando el UUID como clave
		userCopy := user // Crear una copia del usuario para evitar sobrescribir el mismo puntero
		userDB[user.UUID] = &userCopy
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &userDB, nil
}

func (r *cassandra) UpdateUser(ctx context.Context, user *entities.User, UUID string) error {
	existingUser, err := r.GetUserByUUID(ctx, UUID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return r.repository.GetSession().Query(
		"UPDATE users SET username = ?, password = ? WHERE UUID = ?",
		user.Credentials.Username, user.Credentials.PasswordHash, UUID,
	).Exec()
}
