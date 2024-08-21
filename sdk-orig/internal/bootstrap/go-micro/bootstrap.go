/*
La idea seria configurar solo las .envs o variables desde algun lado y listo.
*/

package gmbtrap

import (
	gomicropkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/portspkg"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4/registry"
)

func BootstrapGoMicro() (portspkg.GoMicroService, error) {
	config := gomicropkg.NewGoMicroConfig(
		viper.GetString("APP_NAME"),
		viper.GetString("APP_VERSION"),
		viper.GetString("GOMICRO_MS_PORT"),
	)

	if reg := viper.GetString("CONSUL_ADDRESS"); reg != "" {
		consulReg := consul.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{reg} // reg:CONSUL_ADDRESS=http://consul:8500
		})
		config.SetRegistry(consulReg)
	}

	// if authEnabled := viper.GetBool("GOMICRO_MS_AUTH"); authEnabled {
	// 	authService := auth.NewAuth(
	// 		auth.Credentials(viper.GetString("AUTH_USER"), viper.GetString("AUTH_PASSWORD")),
	// 	)
	// 	config.SetAuth(authService)
	// }

	// if loggerLevel := viper.GetString("LOGGER_LEVEL"); loggerLevel != "" {
	// 	config.SetLogger(logger.NewLogger(
	// 		logger.WithLevel(logger.ParseLevel(loggerLevel)),
	// 	))
	// }
	///////////////////

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

	// Devolver la instancia configurada
	inst, err := gomicropkg.NewGoMicroService(config)
	if err != nil {
		return nil, err
	}

	return inst, nil
}
