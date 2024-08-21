package user

import (
	"context"
	"errors"

	"github.com/gocql/gocql"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql/portspkg"

	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/user/portscore"
)

// cassandraRepository struct con instancia de cliente de Cassandra
type cassandraRepository struct {
	csdInst portspkg.CassandraClient
}

// NewCassandraRepository crea un nuevo repositorio de usuarios en Cassandra
func NewCassandraRepository(inst portspkg.CassandraClient) portscore.Repository {
	return &cassandraRepository{
		csdInst: inst,
	}
}

// SaveUser guarda un nuevo usuario en Cassandra
func (r *cassandraRepository) SaveUser(ctx context.Context, user *entities.User) error {
	return r.csdInst.GetSession().Query(
		"INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.UUID, user.Credentials.Username, user.Credentials.PasswordHash, user.CreatedAt,
	).Exec()
}

// GetUser recupera un usuario por su UUID
func (r *cassandraRepository) GetUser(ctx context.Context, id string) (*entities.User, error) {
	var user entities.User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE id = ?",
		id,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserUUID recupera el UUID de un usuario dado su nombre de usuario y hash de contraseña
func (r *cassandraRepository) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	var uuid string
	err := r.csdInst.GetSession().Query(
		"SELECT id FROM users WHERE username = ? AND password = ?",
		username, passwordHash,
	).Consistency(gocql.One).Scan(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

// GetUserByUsername recupera un usuario por su nombre de usuario
func (r *cassandraRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser elimina un usuario por su UUID
func (r *cassandraRepository) DeleteUser(ctx context.Context, id string) error {
	return r.csdInst.GetSession().Query(
		"DELETE FROM users WHERE id = ?",
		id,
	).Exec()
}

// ListUsers lista todos los usuarios en la base de datos
func (r *cassandraRepository) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	// Crear una instancia de InMemDB
	userDB := make(entities.InMemDB)

	// Crear un iterador para la consulta a la base de datos
	iter := r.csdInst.GetSession().Query("SELECT id, username, password, created_at FROM users").Iter()

	var user entities.User
	for iter.Scan(&user.UUID, &user.Credentials.Username, &user.Credentials.PasswordHash, &user.CreatedAt) {
		// Agregar el usuario al mapa usando el UUID como clave
		userCopy := user // Crear una copia del usuario para evitar sobrescribir el mismo puntero
		userDB[user.UUID] = &userCopy
	}

	// Cerrar el iterador y manejar errores
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &userDB, nil
}

// UpdateUser actualiza la información de un usuario
func (r *cassandraRepository) UpdateUser(ctx context.Context, user *entities.User, id string) error {
	// Verificar si el usuario existe
	existingUser, err := r.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	// Actualizar el usuario
	return r.csdInst.GetSession().Query(
		"UPDATE users SET username = ?, password = ? WHERE id = ?",
		user.Credentials.Username, user.Credentials.PasswordHash, id,
	).Exec()
}
