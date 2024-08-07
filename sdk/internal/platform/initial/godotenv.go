package initialsetup

import (
	"fmt"

	"github.com/joho/godotenv"
)

func SetupDotEnvConfig(configPaths ...string) error {
	for _, path := range configPaths {
		err := godotenv.Load(path)
		if err != nil {
			return fmt.Errorf("error loading .env file from path %s: %w", path, err)
		}
	}
	return nil
}
