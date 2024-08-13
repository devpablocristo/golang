//go:build wireinject
// +build wireinject

package shared

import (
	"github.com/google/wire"

	// authhandler "github.com/devpablocristo/golang/sdk/cmd/gateways/auth"
	// auth "github.com/devpablocristo/golang/sdk/internal/core/auth"

	monhandler "github.com/devpablocristo/golang/sdk/cmd/gateways/monitoring"

	userhandler "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	cass "github.com/devpablocristo/golang/sdk/internal/bootstrap/cassandra"
	user "github.com/devpablocristo/golang/sdk/internal/core/user"
)

func InitializeUserHandler() (*userhandler.GinHandler, error) {
	wire.Build(
		cass.NewCassandraInstance,
		user.NewCassandraRepository,
		user.NewUserUseCases,
		userhandler.NewGinHandler,
	)
	return &userhandler.GinHandler{}, nil
}

// func InitializeAuthHandler() (*authhandler.GinHandler, error) {
// 	wire.Build(
// 		auth.NewAuthUseCases,
// 		authhandler.NewGinHandler,
// 	)
// 	return &authhandler.GinHandler{}, nil
// }

func InitializeMonitoring() (*monhandler.GinHandler, error) {
	wire.Build(
		monhandler.NewGinHandler,
	)
	return &monhandler.GinHandler{}, nil
}
