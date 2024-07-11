//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	hdl "github.com/devpablocristo/qh/events/cmd/rest/handlers"
	core "github.com/devpablocristo/qh/events/internal/core"
	usr "github.com/devpablocristo/qh/events/internal/core/user"
	cass "github.com/devpablocristo/qh/events/internal/platform/cassandra"
	is "github.com/devpablocristo/qh/events/pkg/init-setup"
)

func InitializeUserHandler() (*hdl.UserHandler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		usr.NewUserRepository,
		core.NewUserUseCase,
		hdl.NewUserHandler,
	)
	return &hdl.UserHandler{}, nil
}

func InitializeAuthHandler() (*hdl.AuthHandler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		usr.NewUserRepository,
		core.NewAuthUseCase,
		is.GetJWTSecretKey,
		hdl.NewAuthHandler,
	)
	return &hdl.AuthHandler{}, nil
}
