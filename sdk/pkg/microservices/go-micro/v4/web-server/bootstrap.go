package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
)

func Bootstrap(webRouter interface{}) (ports.WebServer, error) {
	config := newConfigWebServer(
		webRouter,
		viper.GetString("WEB_SERVER_NAME"),
		viper.GetString("CONSUL_ADDRESS"),
		viper.GetString("WEB_SERVER_HOST"),
		viper.GetInt("WEB_SERVER_PORT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newWebServer(config)
}
