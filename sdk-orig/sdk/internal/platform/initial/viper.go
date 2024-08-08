package initialsetup

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetupViperConfig(configPaths ...string) error {
	viper.SetConfigName("./config/.env")
	viper.SetConfigType("env")

	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}
