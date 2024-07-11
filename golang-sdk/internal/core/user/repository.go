package user

import (
	"context"
	"time"

	csdgocsl "github.com/devpablocristo/qh/events/pkg/cassandra/gocql"
	"github.com/gocql/gocql"
)

type repository struct {
	csdInst csdgocsl.CassandraClientPort
}

func NewUserRepository(inst csdgocsl.CassandraClientPort) RepositoryPort {
	return &repository{
		csdInst: inst,
	}
}

func (r *repository) Save(user User) error {
	return r.csdInst.GetSession().Query(
		"INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)",
		user.ID, user.Username, user.Password, user.CreatedAt,
	).Exec()
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	return user, err
}

func (r *repository) GetUser(ctx context.Context, id string) (User, error) {
	return User{
		ID:        id,
		Username:  "sample_user",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
	}, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := r.csdInst.GetSession().Query(
		"SELECT id, username, password, created_at FROM users WHERE username = ?",
		username,
	).Consistency(gocql.One).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	return user, err
}
