package user

import (
	"context"
	"time"

	"github.com/gocql/gocql"

	csdgocsl "github.com/devpablocristo/qh/events/pkg/cassandra/gocql"
)

type repository struct {
	csdInst csdgocsl.CassandraClientPort
}

func NewUserRepository(inst csdgocsl.CassandraClientPort) RepositoryPort {
	return &repository{
		csdInst: inst,
	}
}

func (r *repository) SaveUser(ctx context.Context, user *User) error {
	return r.csdInst.GetSession().Query(
		"INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.UUID, user.Username, user.Password, user.CreatedAt,
	).Exec()
}

func (r *repository) GetUser(ctx context.Context, id string) (*User, error) {
	return &User{
		UUID:      id,
		Username:  "sample_user",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
	}, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.UUID, &user.Username, &user.Password, &user.CreatedAt)
	return &user, err
}
