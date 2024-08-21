package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/user/portscore"
	"github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

// notas implementacion:
// instancia := setup
// repo := repo(instancia)
// usecases =usecases(repo)
// handler:= handler(usecases)
// reg(handler) <-- grpc

type mapDbRepository struct {
	mapDbInst portspkg.MapDbClient
}

// NewMapDbRepository crea un nuevo repositorio de usuarios en memoria
func NewMapDbRepository(inst portspkg.MapDbClient) portscore.Repository {
	return &mapDbRepository{
		mapDbInst: inst,
	}
}

// SaveUser guarda un nuevo usuario en el repositorio
func (r *mapDbRepository) SaveUser(ctx context.Context, user *entities.User) error {
	if user.Credentials.Username == "" {
		return errors.New("username is required")
	}

	// Generar un nuevo UUID para el usuario
	user.UUID = uuid.New().String()
	user.CreatedAt = time.Now()

	// Obtener la base de datos desde la instancia y guardar el usuario
	db := r.mapDbInst.GetDb()

	// Guardar el usuario en la base de datos
	db[user.UUID] = user
	return nil
}

// GetUser obtiene un usuario por su ID (UUID)
func (r *mapDbRepository) GetUser(ctx context.Context, UUID string) (*entities.User, error) {
	db := r.mapDbInst.GetDb()

	// Intentar obtener el usuario del mapa
	user, exists := db[UUID].(*entities.User)
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserUUID obtiene el UUID de un usuario por su nombre de usuario y hash de contraseña
func (r *mapDbRepository) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	db := r.mapDbInst.GetDb()

	for _, v := range db {
		user, ok := v.(*entities.User)
		if ok && user.Credentials.Username == username && user.Credentials.PasswordHash == passwordHash {
			return user.UUID, nil
		}
	}
	return "", errors.New("user not found")
}

// GetUserByUsername obtiene un usuario por su nombre de usuario
func (r *mapDbRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	db := r.mapDbInst.GetDb()

	for _, v := range db {
		user, ok := v.(*entities.User)
		if ok && user.Credentials.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUser elimina un usuario por su ID (UUID)
func (r *mapDbRepository) DeleteUser(ctx context.Context, UUID string) error {
	db := r.mapDbInst.GetDb()

	if _, exists := db[UUID]; !exists {
		return errors.New("user not found")
	}
	delete(db, UUID)
	return nil
}

// ListUsers lista todos los usuarios en el repositorio
func (r *mapDbRepository) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	db := r.mapDbInst.GetDb()

	// Convertir el mapa a *entities.InMemDB
	users := entities.InMemDB{}
	for k, v := range db {
		user, ok := v.(*entities.User)
		if ok {
			users[k] = user
		}
	}
	return &users, nil
}

// UpdateUser actualiza la información de un usuario existente
func (r *mapDbRepository) UpdateUser(ctx context.Context, user *entities.User, UUID string) error {
	db := r.mapDbInst.GetDb()

	existingUser, exists := db[UUID].(*entities.User)
	if !exists {
		return errors.New("user not found")
	}

	if user.Credentials.Username != "" {
		existingUser.Credentials.Username = user.Credentials.Username
	}
	if user.Credentials.PasswordHash != "" {
		existingUser.Credentials.PasswordHash = user.Credentials.PasswordHash
	}
	// Mantener la fecha de creación original
	existingUser.CreatedAt = user.CreatedAt

	db[UUID] = existingUser
	return nil
}
