package initsetup

import (
	"fmt"
)

func InitSetup(configPaths ...string) error {
	if len(configPaths) == 0 {
		configPaths = append(configPaths, ".")
	}

	// if err := SetupDotEnvConfig(configPaths...); err != nil {
	// 	return fmt.Errorf("error setting up Viper: %w", err)
	// }

	if err := SetupViperConfig(configPaths...); err != nil {
		return fmt.Errorf("error setting up Viper: %w", err)
	}

	if err := SetupJWTConfig(); err != nil {
		return fmt.Errorf("error initializing JWT config: %w", err)
	}

	if err := SetupLogger(); err != nil {
		return fmt.Errorf("error setting up Logger: %w", err)
	}

	return nil
}
