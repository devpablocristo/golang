package user

import (
	"context"
	"errors"

	"github.com/gocql/gocql"

	entities "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
	sdkcassandraports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
)

type cassandraRepository struct {
	service sdkcassandraports.Repository
}

func NewCassandraRepository(inst sdkcassandraports.Repository) ports.Repository {
	return &cassandraRepository{
		service: inst,
	}
}

func (r *cassandraRepository) SaveUser(ctx context.Context, user *entities.User) error {
	return r.service.GetSession().Query(
		"INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.UUID, user.Credentials.Username, user.Credentials.PasswordHash, user.CreatedAt,
	).Exec()
}

func (r *cassandraRepository) GetUser(ctx context.Context, id string) (*entities.User, error) {
	var user entities.User
	err := r.service.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE id = ?",
		id,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *cassandraRepository) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	var uuid string
	err := r.service.GetSession().Query(
		"SELECT id FROM users WHERE username = ? AND password = ?",
		username, passwordHash,
	).Consistency(gocql.One).Scan(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (r *cassandraRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := r.service.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *cassandraRepository) DeleteUser(ctx context.Context, id string) error {
	return r.service.GetSession().Query(
		"DELETE FROM users WHERE id = ?",
		id,
	).Exec()
}

func (r *cassandraRepository) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	userDB := make(entities.InMemDB)

	iter := r.service.GetSession().Query("SELECT id, username, password, created_at FROM users").Iter()

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

func (r *cassandraRepository) UpdateUser(ctx context.Context, user *entities.User, id string) error {
	existingUser, err := r.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return r.service.GetSession().Query(
		"UPDATE users SET username = ?, password = ? WHERE id = ?",
		user.Credentials.Username, user.Credentials.PasswordHash, id,
	).Exec()
}
