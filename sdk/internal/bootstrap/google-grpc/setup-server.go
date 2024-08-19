package ggrpcsetup

import (
	//"github.com/spf13/viper"

	ggrpcgpkg "github.com/devpablocristo/golang/sdk/pkg/grpc/google"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
)

func NewGgrpcServerInstance() (portspkg.GgrpcServer, error) {
	// config := ggrpcgpkg.NewGrpcConfig(
	// 	viper.GetString("GRPC_SERVER_HOST"),
	// 	viper.GetInt("GRPC_SERVER_PORT"),
	// 	nil, // Si necesitas TLS, aquí iría la configuración TLS
	// )

	config := ggrpcgpkg.NewGrpcConfig(
		"localhost",
		50051,
		nil,
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := ggrpcgpkg.InitializeGgrpcServer(config); err != nil {
		return nil, err
	}

	return ggrpcgpkg.GetGgrpcServerInstance()
}

func NewGgrpcClientInstance() (portspkg.GgrpcClient, error) {
	// config := ggrpcgpkg.NewGrpcConfig(
	// 	viper.GetString("GRPC_SERVER_HOST"),
	// 	viper.GetInt("GRPC_SERVER_PORT"),
	// 	nil, // Si necesitas TLS, aquí iría la configuración TLS
	// )

	config := ggrpcgpkg.NewGrpcConfig(
		"localhost",
		50051,
		nil,
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := ggrpcgpkg.InitializeGgrpcServer(config); err != nil {
		return nil, err
	}

	return ggrpcgpkg.GetGgrpcClientInstance()
}
