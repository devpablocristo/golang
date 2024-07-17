//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	hdl "github.com/devpablocristo/qh/events/cmd/rest/handlers"
	hnc7 "github.com/devpablocristo/qh/events/cmd/rest/handlers/nimble-cin7"
	core "github.com/devpablocristo/qh/events/internal/core"
	nc7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7"
	cin7 "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/cin7"
	nim "github.com/devpablocristo/qh/events/internal/core/nimble-cin7/nimble"
	usr "github.com/devpablocristo/qh/events/internal/core/user"
	cass "github.com/devpablocristo/qh/events/internal/platform/cassandra"
	rd "github.com/devpablocristo/qh/events/internal/platform/redis"
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

func InitializeCin7NimbleHandler() (*hnc7.Cin7Handler, error) {
	wire.Build(
		rd.NewRedisInstance,
		cin7.NewRedisRepository,
		nc7.NewCin7UseCase,
		hnc7.NewCin7Handler,
	)
	return &hnc7.Cin7Handler{}, nil
}

// integration
// func SetupRoutes(r *gin.Engine) {
//     nimbleRepo := repository.NewNimbleRepository()
//     cin7Repo := cin7repository.NewCin7Repository(config.RedisClient)

//     cin7UseCase := cin7usecase.NewCin7UseCase(cin7Repo)
//     nimbleUseCase := usecase.NewNimbleUseCase(nimbleRepo, cin7UseCase)

//     nimbleHandler := handler.NewNimbleHandler(nimbleUseCase)
//     cin7Handler := cin7handler.NewCin7Handler(cin7UseCase)

//     r.POST("/nimble/orders", nimbleHandler.HandleOrderShipment)
//     r.POST("/cin7/shipments", cin7Handler.HandleShipmentUpdate)
// }
