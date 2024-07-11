package initsetup

import (
	"github.com/spf13/viper"
)

func SetupJWTConfig() error {
	viper.SetDefault("JWT_SECRET_KEY", "secret")
	viper.AutomaticEnv()

	return nil
}

func GetJWTSecretKey() string {
	return viper.GetString("JWT_SECRET_KEY")
}
