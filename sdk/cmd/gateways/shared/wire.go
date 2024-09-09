wireinject
// +build wireinject

package shared
// package shared

// import (
// 	"github.com/google/wire"

// 	jwtsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/jwt"
// 	rabbitmqsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/rabbitmq"

// 	authgin "github.com/devpablocristo/golang/sdk/cmd/gateways/auth"

// 	authucs "github.com/devpablocristo/golang/sdk/internal/core/auth"
// 	// casssetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/cassandra"
// 	// mongin "github.com/devpablocristo/golang/sdk/cmd/gateways/monitoring"
// 	// usergin "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
// 	// userhucs "github.com/devpablocristo/golang/sdk/internal/core/user"
// )

// func InitializeUserHandler() (*userhandler.GinHandler, error) {
// 	wire.Build(
// 		cass.NewCassandraInstance,
// 		user.NewCassandraRepository,
// 		user.NewUserUseCases,
// 		usergin.NewGinHandler,
// 	)
// 	return &userhandler.GinHandler{}, nil
// }

// func InitializeAuthHandler() (*authgin.GinHandler, error) {
// 	wire.Build(
// 		authjwt.NewJWTInstance,
// 		authrmq.NewRabbitMqInstance,
// 		authucs.NewAuthUseCases,
// 		authgin.NewGinHandler,
// 	)
// 	return &authgin.GinHandler{}, nil
// }

// func InitializeMonitoring() (*monhandler.GinHandler, error) {
// 	wire.Build(
// 		monhandler.NewGinHandler,
// 	)
// 	return &monhandler.GinHandler{}, nil
// }

// func InitializeAuthHandler() (*authgin.GinHandler, error) {
// 	wire.Build(
// 		jwtsetup.NewJWTInstance,
// 		rabbitmqsetup.NewRabbitMqInstance,
// 		authucs.NewAuthUseCases,
// 		authgin.NewGinHandler,
// 	)
// 	return &authgin.GinHandler{}, nil
// }
