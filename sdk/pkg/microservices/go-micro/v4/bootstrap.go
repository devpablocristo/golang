package gomicropkg

import (
	"fmt"
	"os"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

// NOTE: La idea es configurar solo las .envs o variables desde algun lado y todo lo demas deberia ser automatico.
func Bootstrap() (ports.Service, error) {
	config := NewConfig(
		viper.GetString("APP_NAME"),
		viper.GetString("APP_VERSION"),
		viper.GetString("GOMICRO_MS_PORT"),
	)

	if reg := viper.GetString("CONSUL_ADDRESS"); reg != "" {
		consulReg := consul.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{reg} // reg:CONSUL_ADDRESS=http://consul:8500
		})
		fmt.Println("consulReg:", consulReg)
		config.SetRegistry(consulReg)
	}

	if authEnabled := viper.GetBool("GOMICRO_MS_AUTH"); authEnabled {
		authService := auth.NewAuth(
			auth.Credentials(
				viper.GetString("AUTHORIZED_USER"),
				viper.GetString("AUTHORIZED_USER_PASSWORD"),
			),
		)
		config.SetAuth(authService)
	}

	if loggerLevel := viper.GetString("LOGGER_LEVEL"); loggerLevel != "" {
		loggerService := logger.NewLogger(
			logger.WithLevel(logger.InfoLevel),
			logger.WithOutput(os.Stdout),
		)
		config.SetLogger(loggerService)
	}

	// Configurar el broker si está presente
	// if brokerAddress := viper.GetString("BROKER_ADDRESS"); brokerAddress != "" {
	// 	config.SetBroker(broker.NewBroker(
	// 		broker.Addrs(brokerAddress),
	// 	))
	// }

	// // Configurar el cliente si está presente
	// if clientTimeout := viper.GetDuration("CLIENT_TIMEOUT"); clientTimeout > 0 {
	// 	config.SetClient(client.NewClient(
	// 		client.RequestTimeout(clientTimeout),
	// 	))
	// }

	// // Configurar el servidor si está presente
	// if serverAddress := viper.GetString("SERVER_ADDRESS"); serverAddress != "" {
	// 	config.SetServer(server.NewServer(
	// 		server.Address(serverAddress),
	// 	))
	// }

	// // Configurar el almacenamiento si está presente
	// if storeType := viper.GetString("STORE_TYPE"); storeType != "" {
	// 	config.SetStore(store.NewStore(
	// 		store.WithBackend(storeType),
	// 	))
	// }

	// // Configurar el transporte si está presente
	// if transportType := viper.GetString("TRANSPORT_TYPE"); transportType != "" {
	// 	config.SetTransport(transport.NewTransport(
	// 		transport.WithOption(transportType),
	// 	))
	// }

	// // Configurar el servicio web si está presente
	// if webAddress := viper.GetString("WEB_ADDRESS"); webAddress != "" {
	// 	config.SetWebService(web.NewService(
	// 		web.Address(webAddress),
	// 	))
	// }

	return NewService(config)
}