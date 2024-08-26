package sdkviper

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(configPaths ...string) error {
	if len(configPaths) == 0 {
		configPaths = append(configPaths, ".")
	}

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
