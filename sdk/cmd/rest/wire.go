//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	hauth "github.com/devpablocristo/qh/events/cmd/rest/auth/handlers"
	hnc7 "github.com/devpablocristo/qh/events/cmd/rest/nimble-cin7/handlers"
	husr "github.com/devpablocristo/qh/events/cmd/rest/user/handlers"
	core "github.com/devpablocristo/qh/events/internal/core"
	nc7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7"
	cin7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/cin7"
	nim "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/nimble"
	usr "github.com/devpablocristo/qh/events/internal/core/user"
	cass "github.com/devpablocristo/qh/events/internal/platform/cassandra"
	rd "github.com/devpablocristo/qh/events/internal/platform/redis"
	is "github.com/devpablocristo/qh/events/pkg/init-setup"  
)

func InitializeUserHandler() (*husr.UserHandler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		usr.NewUserRepository,
		core.NewUserUseCase,
		husr.NewUserHandler,
	)
	return &husr.UserHandler{}, nil
}

func InitializeAuthHandler() (*hauth.AuthHandler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		usr.NewUserRepository,
		core.NewAuthUseCase,
		is.GetJWTSecretKey,
		hauth.NewAuthHandler,
	)
	return &hauth.AuthHandler{}, nil
}

func InitializeNimbleHandler() (*hnc7.NimbleHandler, error) {
	wire.Build(
		rd.NewRedisInstance,
		nim.NewRedisRepository,
		cin7.NewRedisRepository,
		nc7.NewCin7UseCase,
		nc7.NewNimbleUseCase,
		hnc7.NewNimbleHandler,
	)
	return &hnc7.NimbleHandler{}, nil
}

func InitializeCin7Handler() (*hnc7.Cin7Handler, error) {
	wire.Build(
		rd.NewRedisInstance,
		cin7.NewRedisRepository,
		nc7.NewCin7UseCase,
		hnc7.NewCin7Handler,
	)
	return &hnc7.Cin7Handler{}, nil
}
