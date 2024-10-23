// wire.go
package main

import (
	"github.com/google/wire"

	authconn "github.com/devpablocristo/golang/sdk/sg/auth/internal/adapters/connectors"
	authgtw "github.com/devpablocristo/golang/sdk/sg/auth/internal/adapters/gateways"
	auth "github.com/devpablocristo/golang/sdk/sg/auth/internal/core"
)

// ProviderSet contains all the providers needed for the application
var ProviderSet = wire.NewSet(
	authconn.NewJwtService,
	authconn.NewHttpClient,
	authconn.NewGorillaSessionManager,
	authconn.NewPostgreSQL,
	auth.NewUseCases,
	authgtw.NewGinHandler,
)

// Application represents the complete application with all its dependencies
type Application struct {
	Handler *authgtw.GinHandler
}

// NewApplication creates a new application instance with all dependencies wired
func NewApplication() (*Application, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Application), "*"),
	)
	return nil, nil
}

// go generate -v -tags=wireinject ./...
