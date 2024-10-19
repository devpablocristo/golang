package sdkhclnt

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"
)

func Bootstrap(tokenEndPointEnvName, clientIDEnvName, clientSecretEnvName, addParamsEnvName string) (ports.Client, ports.Config, error) {
	tokenEndPoint := viper.GetString(tokenEndPointEnvName)
	if tokenEndPoint == "" {
		return nil, nil, fmt.Errorf("token endpoint is empty. Check if %s environment variable is set", tokenEndPointEnvName)
	}

	config := newConfig(
		tokenEndPoint,
		viper.GetString(clientIDEnvName),
		viper.GetString(clientSecretEnvName),
		viper.GetStringMapString(addParamsEnvName),
	)

	if err := config.Validate(); err != nil {
		return nil, nil, err
	}

	client, err := newClient(config)
	if err != nil {
		return nil, nil, err
	}

	return client, config, nil
}
