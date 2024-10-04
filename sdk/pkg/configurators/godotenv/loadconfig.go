package sdkgodotenv

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadConfig carga múltiples archivos .env desde las rutas proporcionadas.
// Si no se especifican rutas, carga el archivo .env del directorio actual.
// Retorna un error en caso de que no se pueda cargar algún archivo .env.
func LoadConfig(configPaths ...string) error {
	// Si no se proporcionan rutas, usa el directorio actual.
	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}

	// Cargar cada archivo .env desde las rutas especificadas.
	for _, path := range configPaths {
		envFilePath := filepath.Join(path, ".env")
		fmt.Printf("Attempting to load .env file from: %s\n", envFilePath) // Debugging
		if err := godotenv.Load(envFilePath); err != nil {
			return fmt.Errorf("failed to load .env file from path %s: %w", envFilePath, err)
		}
	}

	fmt.Println("Environment variables loaded successfully from .env files")
	return nil
}
