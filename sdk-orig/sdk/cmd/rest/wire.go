//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	//hauth "github.com/devpablocristo/golang/sdk/cmd/rest/auth/handlers"
	//cin7 "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/cin7/handler"
	//nimble "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/nimble/handler"
	monitoring "github.com/devpablocristo/golang/sdk/cmd/rest/monitoring/handler"
	userhandler "github.com/devpablocristo/golang/sdk/cmd/rest/user/handler"
	"github.com/devpablocristo/golang/sdk/internal/core"

	//nc7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7"
	//nim "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
	"github.com/devpablocristo/golang/sdk/internal/core/user"
	cass "github.com/devpablocristo/golang/sdk/internal/platform/cassandra"
	//rd "github.com/devpablocristo/golang/sdk/internal/platform/redis"
	//is "github.com/devpablocristo/golang/sdk/pkg/init-setup"
)

func InitializeUserHandler() (*userhandler.Handler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		user.NewUserRepository,
		core.NewUserUseCases,
		userhandler.NewHandler,
	)
	return &userhandler.Handler{}, nil
}

func InitializeMonitoring() (*monitoring.Handler, error) {
	wire.Build(
		monitoring.NewHandler,
	)
	return &monitoring.Handler{}, nil
}

// func InitializeAuthHandler() (*hauth.AuthHandler, error) {
// 	wire.Build(
// 		cass.NewCassandraInstance,
// 		usr.NewUserRepository,
// 		core.NewAuthUseCases,
// 		is.GetJWTSecretKey,
// 		hauth.NewAuthHandler,
// 	)
// 	return &hauth.AuthHandler{}, nil
// }

// func InitializeNimbleHandler() (*nimble.Handler, error) {
// 	wire.Build(
// 		rd.NewRedisInstance,
// 		nim.NewRedisRepository,
// 		cin7.NewRedisRepository,
// 		nc7.NewCin7UseCases,
// 		nc7.NewNimbleUseCases,
// 		cin7.NewCin7Handler,
// 	)
// 	return &hnc7.NimbleHandler{}, nil
// }

// func InitializeCin7Handler() (*hnc7.Cin7Handler, error) {
// 	wire.Build(
// 		rd.NewRedisInstance,
// 		cin7.NewRedisRepository,
// 		nc7.NewCin7UseCases,
// 		hnc7.NewCin7Handler,
// 	)
// 	return &hnc7.Cin7Handler{}, nil
// }
