package sdkjwt

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/defs"
)

func Bootstrap() (defs.Service, error) {
	config := newConfig(
		viper.GetString("JWT_SECRET_KEY"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
