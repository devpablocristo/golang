package sdkviper

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig carga las configuraciones desde archivos de configuración (.yaml, .json, .toml, etc.)
// y archivos de variables de entorno (.env). Si no se especifican rutas, usa el directorio actual.
// Se prioriza el archivo de configuración si existe, y Viper también cargará las variables de entorno.
func LoadConfig(configPaths ...string) error {
	// Paso 1: Configurar Viper para leer las variables de entorno
	configureViper()

	// Si no se especifican rutas, usa el directorio actual.
	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}
	// Añadir rutas donde Viper buscará los archivos de configuración.
	for _, path := range configPaths {
		fmt.Printf("Searching for config files in: %s\n", path) // Debugging
		viper.AddConfigPath(path)
	}

	// Paso 2: Cargar configuraciones desde archivos de múltiples formatos
	err := loadConfigurationFiles()
	if err != nil {
		return err
	}

	fmt.Println("Configuration loaded successfully.")
	return nil
}

// configureViper configura Viper para leer variables de entorno automáticamente.
func configureViper() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Reemplaza puntos por guiones bajos si es necesario
}

// loadConfigurationFiles carga archivos de configuración en los formatos soportados (yaml, json, toml).
// Si no encuentra ningún archivo, genera un error.
func loadConfigurationFiles() error {
	supportedFormats := []string{"yaml", "json", "toml"} // Formatos soportados
	viper.SetConfigName("config")                        // Nombre base para los archivos de configuración

	// Intentar leer cada formato
	for _, format := range supportedFormats {
		viper.SetConfigType(format)
		err := viper.ReadInConfig()
		if err == nil {
			fmt.Printf("Configuration successfully loaded from %s file\n", viper.ConfigFileUsed())
			return nil
		}

		// Si el archivo no se encuentra, continuar probando los otros formatos
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading %s config file: %w", format, err)
		}
	}

	// Si llegamos aquí, significa que no se encontró ningún archivo de configuración
	return fmt.Errorf("no configuration file found in supported formats (yaml, json, toml)")
}

// UnmarshalConfig mapea la configuración cargada a una estructura proporcionada por el usuario.
// Retorna un error si no puede deserializar la configuración.
func UnmarshalConfig(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode configuration into struct: %w", err)
	}
	return nil
}
