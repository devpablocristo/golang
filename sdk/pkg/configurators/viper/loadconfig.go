package sdkviper

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig carga las configuraciones desde archivos de configuración (.yaml, .yml, .json, .toml, .ini, .env).
// Si no se especifican rutas, usa el directorio actual.
// Se prioriza el archivo de configuración si existe, y Viper también cargará las variables de entorno.
func LoadConfig(configPaths ...string) error {
	// Paso 1: Configurar Viper para leer las variables de entorno
	configureViper()

	// Si no se especifican rutas, usa el directorio actual.
	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}

	// Añadir rutas donde Viper buscará los archivos de configuración
	for _, configPath := range configPaths {
		fmt.Printf("Searching for config files: %s\n", configPath) // Debugging

		// Extraer el nombre del archivo y la extensión
		fileName := filepath.Base(configPath)
		fileExtension := filepath.Ext(fileName)

		if fileExtension == "" {
			return fmt.Errorf("the file path must contain a file name with extension")
		}

		// Remover el punto (.) de la extensión para utilizarlo en viper
		fileExtension = strings.TrimPrefix(fileExtension, ".")
		fileNameWithoutExt := strings.TrimSuffix(fileName, "."+fileExtension)

		// Configurar Viper con el nombre del archivo y el tipo
		viper.AddConfigPath(filepath.Dir(configPath))
		viper.SetConfigName(fileNameWithoutExt)
		viper.SetConfigType(fileExtension)

		// Intentar leer el archivo de configuración
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file (%s): %w", configPath, err)
		}

		fmt.Printf("Configuration successfully loaded from %s\n", viper.ConfigFileUsed())
	}

	fmt.Println("Configuration loaded successfully.")
	return nil
}

// configureViper configura Viper para leer variables de entorno automáticamente.
func configureViper() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Reemplaza puntos por guiones bajos si es necesario
}

// UnmarshalConfig mapea la configuración cargada a una estructura proporcionada por el usuario.
// Retorna un error si no puede deserializar la configuración.
func UnmarshalConfig(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode configuration into struct: %w", err)
	}
	return nil
}
