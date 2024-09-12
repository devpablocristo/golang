package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	entities "github.com/devpablocristo/golang/sdk/examples/users-api/internal/user/entities"
	ports "github.com/devpablocristo/golang/sdk/examples/users-api/internal/user/ports"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
)

type mapDbRepository struct {
	service sdkports.Repository
}

func NewMapDbRepository(s sdkports.Repository) ports.Repository {
	return &mapDbRepository{
		service: s,
	}
}

// SaveUser guarda un nuevo usuario en el repositorio
func (r *mapDbRepository) SaveUser(ctx context.Context, user *entities.User) error {
	if user.Credentials.Username == "" {
		return errors.New("username is required")
	}

	user.UUID = uuid.New().String()
	user.CreatedAt = time.Now()

	db := r.service.GetDb()

	db[user.UUID] = user
	return nil
}

func (r *mapDbRepository) GetUser(ctx context.Context, UUID string) (*entities.User, error) {
	db := r.service.GetDb()

	user, exists := db[UUID].(*entities.User)
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *mapDbRepository) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	db := r.service.GetDb()

	for _, v := range db {
		user, ok := v.(*entities.User)
		if ok && user.Credentials.Username == username && user.Credentials.PasswordHash == passwordHash {
			return user.UUID, nil
		}
	}
	return "", errors.New("user not found")
}

func (r *mapDbRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	db := r.service.GetDb()

	for _, v := range db {
		user, ok := v.(*entities.User)
		if ok && user.Credentials.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *mapDbRepository) DeleteUser(ctx context.Context, UUID string) error {
	db := r.service.GetDb()

	if _, exists := db[UUID]; !exists {
		return errors.New("user not found")
	}
	delete(db, UUID)
	return nil
}

func (r *mapDbRepository) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	db := r.service.GetDb()

	users := entities.InMemDB{}
	for k, v := range db {
		user, ok := v.(*entities.User)
		if ok {
			users[k] = user
		}
	}
	return &users, nil
}

func (r *mapDbRepository) UpdateUser(ctx context.Context, user *entities.User, UUID string) error {
	db := r.service.GetDb()

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

	existingUser.CreatedAt = user.CreatedAt

	db[UUID] = existingUser
	return nil
}
