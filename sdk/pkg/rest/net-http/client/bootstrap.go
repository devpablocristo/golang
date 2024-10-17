package sdkhclnt

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"
)

func Bootstrap() (ports.Client, ports.Config, error) {
	config := newConfig(
		viper.GetString("AUTH_SERVER_URL"),
		viper.GetString("REALM"),
		viper.GetString("CLIENT_ID"),
		viper.GetString("CLIENT_SECRET"),
	)

	if err := config.Validate(); err != nil {
		return nil, nil, err
	}

	client, err := newClient(config)
	if err != nil {
		return nil, nil, err
	}

	return client, config, err
}
