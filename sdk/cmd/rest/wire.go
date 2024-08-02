//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	hauth "github.com/devpablocristo/golang/sdk/cmd/rest/auth/handlers"
	cin7 "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/cin7/handler"
	nimble "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/nimble/handler"
	restuser "github.com/devpablocristo/golang/sdk/cmd/rest/user/handler"
	core "github.com/devpablocristo/golang/sdk/internal/core"
	nc7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7"
	nim "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
	usr "github.com/devpablocristo/golang/sdk/internal/core/user"
	cass "github.com/devpablocristo/golang/sdk/internal/platform/cassandra"
	rd "github.com/devpablocristo/golang/sdk/internal/platform/redis"
	is "github.com/devpablocristo/golang/sdk/pkg/init-setup"
)

func InitializeUserHandler() (*restuser.Handler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		usr.NewUserRepository,
		core.NewUserUseCase,
		restuser.NewHandler,
	)
	return &restuser.Handler{}, nil
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

func InitializeNimbleHandler() (*nimble.Handler, error) {
	wire.Build(
		rd.NewRedisInstance,
		nim.NewRedisRepository,
		cin7.NewRedisRepository,
		nc7.NewCin7UseCase,
		nc7.NewNimbleUseCase,
		cin7.NewCin7Handler,
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
