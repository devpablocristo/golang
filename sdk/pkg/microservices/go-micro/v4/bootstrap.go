package gomicropkg

import (
	"fmt"
	"os"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func Bootstrap() (ports.Service, error) {
	config := NewConfig()

	if err := setupService(config); err != nil {
		return nil, err
	}

	if err := setupRegistry(config); err != nil {
		return nil, err
	}

	if err := setupAuth(config); err != nil {
		return nil, err
	}

	if err := setupLogger(config); err != nil {
		return nil, err
	}

	if err := setupWebService(config); err != nil {
		return nil, err
	}

	return NewService(config)
}

func setupService(config ports.Config) error {
	appName := viper.GetString("APP_NAME")
	appVersion := viper.GetString("APP_VERSION")
	msPort := viper.GetString("GOMICRO_MS_PORT")

	if appName == "" || appVersion == "" || msPort == "" {
		return fmt.Errorf("missing required service configuration: APP_NAME, APP_VERSION, or GOMICRO_MS_PORT")
	}

	service := micro.NewService(
		micro.Name(appName),
		micro.Version(appVersion),
		micro.Address(msPort),
	)

	config.SetService(service)
	return nil
}

func setupRegistry(config ports.Config) error {
	consulAddress := viper.GetString("CONSUL_ADDRESS")
	if consulAddress == "" {
		return nil // Consul no es obligatorio, continuar sin error
	}

	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{consulAddress}
	})

	config.SetRegistry(consulReg)
	return nil
}

func setupAuth(config ports.Config) error {
	if !viper.GetBool("GOMICRO_MS_AUTH") {
		return nil // Autenticación no habilitada, continuar sin error
	}

	authUser := viper.GetString("AUTHORIZED_USER")
	authPassword := viper.GetString("AUTHORIZED_USER_PASSWORD")

	if authUser == "" || authPassword == "" {
		return fmt.Errorf("missing required authentication configuration: AUTHORIZED_USER or AUTHORIZED_USER_PASSWORD")
	}

	authService := auth.NewAuth(
		auth.Credentials(authUser, authPassword),
	)

	config.SetAuth(authService)
	return nil
}

func setupLogger(config ports.Config) error {
	loggerLevel := viper.GetString("LOGGER_LEVEL")
	if loggerLevel == "" {
		return nil // Logger no configurado, continuar sin error
	}

	loggerService := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel), // Cambia esto si usas otros niveles
		logger.WithOutput(os.Stdout),
	)

	config.SetLogger(loggerService)
	return nil
}

func setupWebService(config ports.Config) error {
	webAddress := viper.GetString("ROUTER_PORT")
	if webAddress == "" {
		return nil // WebService no configurado, continuar sin error
	}

	webService := web.NewService(
		web.Name(viper.GetString("APP_NAME")+"-web"),
		web.Version(viper.GetString("APP_VERSION")),
		web.Address(":"+webAddress),
	)

	config.SetWebService(webService)
	return nil
}
