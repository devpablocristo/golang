// wire.go
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/wire"

	// sdkaws "github.com/devpablocristo/golang/sdk/pkg/aws/localstack"
	// sdkawsports "github.com/devpablocristo/golang/sdk/pkg/aws/localstack/ports"

	authconn "github.com/devpablocristo/golang/sdk/sg/auth/internal/adapters/connectors"
	authgtw "github.com/devpablocristo/golang/sdk/sg/auth/internal/adapters/gateways"

	auth "github.com/devpablocristo/golang/sdk/sg/auth/internal/core"
	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

// ProvideAWSStack proporciona la configuración de AWS (por ejemplo, LocalStack)
// func ProvideAWSStack() (sdkawsports.Stack, error) {
// 	return sdkaws.Bootstrap()
// }

// // ProvideS3Client crea un cliente de S3 utilizando la configuración de AWS
// func ProvideS3Client(stack sdkawsports.Stack) *s3.Client {
// 	return s3.NewFromConfig(stack.GetCfg())
// }

// ProvideAuthUsecases crea la capa de casos de uso de autenticación
func ProvideAuthUsecases(
	jwtService ports.JwtService,
	repository ports.Repository,
	httpClient ports.HttpClient,
	sessionManager ports.SessionManager) ports.UseCases {
	return auth.NewUseCases(jwtService, repository, httpClient, sessionManager)
}

// ProvideGinHandler inicializa el manejador Gin con los casos de uso de autenticación
func ProvideGinHandler(usecases ports.UseCases) (*authgtw.GinHandler, error) {
	return authgtw.NewGinHandler(usecases)
}

// Injector es el que ensamblará todas las dependencias
func InitializeApplication() (*authgtw.GinHandler, error) {
	wire.Build(
		// ProvideAWSStack,
		// ProvideS3Client,
		authconn.NewJwtService,
		authconn.NewHttpClient,
		authconn.NewGorillaSessionManager,
		authconn.NewPostgreSQL,
		ProvideAuthUsecases,
		ProvideGinHandler,
	)
	return nil, nil
}

// go generate -v -tags=wireinject ./...
