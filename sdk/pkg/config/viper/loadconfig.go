package sdkviper

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	sdkgodotenv "github.com/devpablocristo/golang/sdk/pkg/config/godotenv"
	sdktools "github.com/devpablocristo/golang/sdk/pkg/tools"
)

// LoadConfig carga archivos de configuración compatibles con Viper y .env
func LoadConfig(configPaths ...string) error {
	configureViper()

	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}

	// Buscar los archivos en el proyecto utilizando FilesFinder
	foundFiles, err := sdktools.FilesFinder(configPaths...)
	if err != nil {
		return fmt.Errorf("SDK Viper: FilesFinder failed to find files: %w", err)
	}

	configDirs := make(map[string]bool)
	for _, configPath := range foundFiles {
		// Cargar archivos .env con godotenv
		if isEnvFile(configPath) {
			if err := sdkgodotenv.LoadConfig(configPath); err != nil {
				fmt.Printf("SDK Viper: Failed to load .env file '%s': %v\n", configPath, err)
				continue
			}
			fmt.Printf("SDK Viper: Successfully loaded .env file from %s\n", configPath)
			continue
		}

		// Cargar otros archivos con Viper
		if err := loadViperConfig(configPath, configDirs); err != nil {
			fmt.Printf("SDK Viper: %v\n", err)
			continue
		}
	}

	return nil
}

// isEnvFile verifica si el archivo es un .env
func isEnvFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".env")
}

// loadViperConfig carga un archivo de configuración con Viper
func loadViperConfig(configPath string, configDirs map[string]bool) error {
	fileNameWithoutExt, fileExtension, err := sdktools.FileNameAndExtension(configPath)
	if err != nil {
		return fmt.Errorf("SDK Viper: Skipping file '%s': %v", configPath, err)
	}

	viper.SetConfigName(fileNameWithoutExt)
	viper.SetConfigType(fileExtension)

	dir := filepath.Dir(configPath)
	if !configDirs[dir] {
		viper.AddConfigPath(dir)
		configDirs[dir] = true
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("SDK Viper: No config files found at dir '%s'", dir)
		}
		return fmt.Errorf("SDK Viper: Skipping invalid config file '%s': %v", configPath, err)
	}

	fmt.Printf("SDK Viper: Configuration successfully loaded from %s\n", viper.ConfigFileUsed())
	return nil
}

// configureViper configura Viper para cargar variables de entorno
func configureViper() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// UnmarshalConfig decodifica la configuración en una estructura Go
func UnmarshalConfig(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode configuration into struct: %w", err)
	}
	return nil
}
