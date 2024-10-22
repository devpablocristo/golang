// wire.go
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	userconn "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors"
	usergtw "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/gateways"
	companyconn "github.com/devpablocristo/golang/sdk/sg/users/internal/company/adapters/connectors"
	personconn "github.com/devpablocristo/golang/sdk/sg/users/internal/person/adapters/connectors"

	companyports "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/ports"
	user "github.com/devpablocristo/golang/sdk/sg/users/internal/core"
	userports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
	personports "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/ports"
)

// ProvideAuthUsecases crea la capa de casos de uso de autenticación
func ProvideUserUsecases(
	userRepo userports.Repository,
	personRepo personports.Repository,
	companyRepo companyports.Repository) userports.UseCases {
	return user.NewUseCases(userRepo, personRepo, companyRepo)
}

// ProvideGinHandler inicializa el manejador Gin con los casos de uso de autenticación
func ProvideUserGinHandler(usecases userports.UseCases) (*usergtw.GinHandler, error) {
	return usergtw.NewGinHandler(usecases)
}

// Injector es el que ensamblará todas las dependencias
func InitializeApplication() (*usergtw.GinHandler, error) {
	wire.Build(
		userconn.NewPostgreSQL,
		personconn.NewPostgreSQL,
		companyconn.NewPostgreSQL,
		ProvideUserUsecases,
		ProvideUserGinHandler,
	)
	return nil, nil
}

// go generate -v -tags=wireinject ./...
