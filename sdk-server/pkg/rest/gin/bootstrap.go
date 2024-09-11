package sdkgin

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

func Bootstrap() (ports.Server, error) {
	config := newConfig(viper.GetString("ROUTER_PORT"))

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newServer(config)
}
