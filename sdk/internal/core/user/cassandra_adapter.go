package user

import (
	"context"
	"errors"

	"github.com/gocql/gocql"

	csdgocsl "github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql"
)

// repository struct con instancia de cliente de Cassandra
type repository struct {
	csdInst csdgocsl.CassandraClientPort
}

// NewUserRepository crea un nuevo repositorio de usuarios
func NewUserRepository(inst csdgocsl.CassandraClientPort) Repository {
	return &repository{
		csdInst: inst,
	}
}

// SaveUser guarda un nuevo usuario en Cassandra
func (r *repository) SaveUser(ctx context.Context, user *User) error {
	return r.csdInst.GetSession().Query(
		"INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.UUID, user.Username, user.PasswordHash, user.CreatedAt,
	).Exec()
}

// GetUser recupera un usuario por su UUID
func (r *repository) GetUser(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE id = ?",
		id,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername recupera un usuario por su nombre de usuario
func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser elimina un usuario por su UUID
func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return r.csdInst.GetSession().Query(
		"DELETE FROM users WHERE id = ?",
		id,
	).Exec()
}

// ListUsers lista todos los usuarios
// ListUsers lista todos los usuarios en la base de datos
func (r *repository) ListUsers(ctx context.Context) (*InMemDB, error) {
	// Crear una instancia de InMemDB
	userDB := make(InMemDB)

	// Crear un iterador para la consulta a la base de datos
	iter := r.csdInst.GetSession().Query("SELECT id, username, password, created_at FROM users").Iter()

	var user User
	for iter.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.CreatedAt) {
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
func (r *repository) UpdateUser(ctx context.Context, user *User, id string) error {
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
		user.Username, user.PasswordHash, id,
	).Exec()
}
